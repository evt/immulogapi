package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/evt/immulogapi/internal/app/interceptors"
	"github.com/evt/immulogapi/internal/app/services/authservice"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/evt/immulogapi/internal/app/config"
	"github.com/evt/immulogapi/internal/app/repo"
	"github.com/evt/immulogapi/internal/app/services/logservice"
	"github.com/evt/immulogapi/internal/pkg/immudb"
	"github.com/evt/immulogapi/internal/pkg/jwt"
	v1 "github.com/evt/immulogapi/proto/v1"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	log.Println("reading config")
	cnf := config.Init()

	ctx := context.Background()

	grpcAddr := fmt.Sprintf("%s:%d", cnf.App.GrpcHost, cnf.App.GrpcPort)
	httpAddr := fmt.Sprintf("%s:%d", cnf.App.HttpHost, cnf.App.HttpPort)
	immudbAddr := fmt.Sprintf("%s:%d", cnf.ImmuDB.Host, cnf.ImmuDB.Port)

	// connect to immudb
	log.Printf("connecting to immudb at %s", immudbAddr)

	db, err := immudb.New(cnf)
	if err != nil {
		return fmt.Errorf("failed connecting to immudb: %w", err)
	}
	defer db.Close()

	// create log repo
	logRepo := repo.NewImmuLogRepo(db)
	// create log table if not exists
	if err := logRepo.CreateLogLinesTable(ctx); err != nil {
		return fmt.Errorf("failed creating loglines table in immudb: %w", err)
	}
	// create user repo
	userRepo := repo.NewImmuUserRepo(db)
	// create users table if not exists
	if err := userRepo.CreateUsersTable(ctx); err != nil {
		return fmt.Errorf("failed creating users table in immudb: %w", err)
	}
	// create test user if not exists
	if err := userRepo.CreateTestUser(ctx); err != nil {
		return fmt.Errorf("failed creating test user: %w", err)
	}

	// create JWT manager
	jwtManager := jwt.NewManager(cnf.App.JWTSecret)

	// create log service
	logService := logservice.New(logRepo)
	// create user service
	authService := authservice.New(userRepo, jwtManager)

	// create gRPC listener
	grpcListener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return fmt.Errorf("failed to start grpc listener: %w", err)
	}
	defer grpcListener.Close()

	// create auth unary interceptor
	authInterceptor := interceptors.NewAuthInterceptor(jwtManager)

	// create grpc server with auth interceptor
	opts := append(grpcServerOptions(), grpc.UnaryInterceptor(authInterceptor.Unary()))
	grpcServer := grpc.NewServer(opts...)

	v1.RegisterLogServiceServer(grpcServer, logService)
	v1.RegisterAuthServiceServer(grpcServer, authService)

	// connect to grpc server for grpc gateway
	grpcCtx, grpcCancel := context.WithTimeout(ctx, time.Minute)
	defer grpcCancel()

	var grpcDialOpts = append(
		grpcClientOptions(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(1024*1024*10), // 10 MB
			grpc.MaxCallSendMsgSize(1024*1024*10), // 10 MB
		),
	)

	conn, err := grpc.DialContext(grpcCtx, grpcAddr, grpcDialOpts...)
	if err != nil {
		return fmt.Errorf("failed connecting to grpc server: %w", err)
	}

	mux := runtime.NewServeMux()
	if err = v1.RegisterLogServiceHandler(ctx, mux, conn); err != nil {
		return fmt.Errorf("v1.RegisterLogServiceHandler failed: %w", err)
	}
	if err = v1.RegisterAuthServiceHandler(ctx, mux, conn); err != nil {
		return fmt.Errorf("v1.RegisterAuthServiceHandler failed: %w", err)
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var errc = make(chan error, 1)
	var wg sync.WaitGroup
	var wgDone = make(chan bool)

	// create HTTP gateway listener
	httpListener, err := net.Listen("tcp", httpAddr)
	if err != nil {
		return fmt.Errorf("failed to start http listener: %w", err)
	}
	defer httpListener.Close()

	wg.Add(2)

	go func() {
		defer wg.Done()
		log.Printf("running gRPC server on %s", grpcAddr)
		if err := grpcServer.Serve(grpcListener); err != nil {
			errc <- fmt.Errorf("grpcServer.Serve failed: %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		log.Printf("running HTTP server on %s", httpAddr)
		if err := http.Serve(httpListener, mux); err != nil {
			errc <- fmt.Errorf("grpcServer.Serve failed: %v", err)
		}
	}()

	go func() {
		wg.Wait()
		close(wgDone)
	}()

	select {
	case <-wgDone:
		break
	case err := <-errc:
		return err
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		// context cancelled
	case sig := <-c:
		// signal received
		log.Printf("received signal: %v\n", sig)
		cancel()
	}

	grpcServer.GracefulStop()

	log.Println("gRPC server stopped gracefully")

	return nil
}

func grpcServerOptions() []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			otelgrpc.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			otelgrpc.StreamServerInterceptor(),
		),
	}
}

func grpcClientOptions() []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithChainUnaryInterceptor(
			otelgrpc.UnaryClientInterceptor(),
		),
		grpc.WithChainStreamInterceptor(
			otelgrpc.StreamClientInterceptor(),
		),
	}
}
