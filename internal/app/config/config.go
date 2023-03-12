package config

import (
	"log"
	"os"
	"strconv"
)

// Config is a config
type Config struct {
	ImmuDB ImmuDBConfig
	App    AppConfig
}

type ImmuDBConfig struct {
	Username string
	Password string
	Database string
	Port     int
	Host     string
}

type AppConfig struct {
	HttpHost  string
	HttpPort  int
	GrpcHost  string
	GrpcPort  int
	JWTSecret string
}

// Init reads config from environment.
func Init() *Config {
	immudbPort, err := strconv.Atoi(os.Getenv("IMMUDB_PORT"))
	if err != nil {
		log.Fatalf("IMMUDB_PORT doesn't look like an integer: %s", err)
	}
	httpPort, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		log.Fatalf("HTTP_PORT doesn't look like an integer: %s", err)
	}
	grpcPort, err := strconv.Atoi(os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalf("GRPC_PORT doesn't look like an integer: %s", err)
	}
	return &Config{
		ImmuDB: ImmuDBConfig{
			Username: os.Getenv("IMMUDB_USERNAME"),
			Password: os.Getenv("IMMUDB_PASSWORD"),
			Database: os.Getenv("IMMUDB_DATABASE"),
			Host:     os.Getenv("IMMUDB_HOST"),
			Port:     immudbPort,
		},
		App: AppConfig{
			HttpHost: os.Getenv("HTTP_HOST"),
			HttpPort: httpPort,
			GrpcHost: os.Getenv("GRPC_HOST"),
			GrpcPort: grpcPort,
		},
	}
}
