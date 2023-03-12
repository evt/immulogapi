package authservice

import (
	"context"
	"log"

	"github.com/evt/immulogapi/internal/app/models"
	"github.com/evt/immulogapi/internal/pkg/jwt"
	v1 "github.com/evt/immulogapi/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	v1.UnimplementedAuthServiceServer
	signer   Signer
	userRepo UserRepo
}

func New(userRepo UserRepo, signer Signer) *Service {
	return &Service{
		userRepo: userRepo,
		signer:   signer,
	}
}

func (s *Service) Auth(ctx context.Context, r *v1.AuthRequest) (*v1.AuthResponse, error) {
	user, err := s.userRepo.GetUser(ctx, r.Login, r.Pass)
	if err != nil {
		log.Printf("failed to get user from repository: %v", err)
		return nil, status.Error(codes.Internal, "failed to get user from repository")
	}

	token, err := s.signer.Sign(&jwt.CustomClaims{
		UserID: user.ID,
	})
	if err != nil {
		log.Printf("failed signing claims: %v", err)
		return nil, status.Error(codes.Internal, "failed signing claims")
	}

	return &v1.AuthResponse{Token: token}, nil
}

func (s *Service) CreateUser(ctx context.Context, r *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	user := &models.User{
		Login: r.Login,
		Pass:  r.Pass,
	}

	userID, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		log.Printf("failed to create user in repository: %v", err)
		return nil, status.Error(codes.Internal, "failed to create user in repository")
	}

	return &v1.CreateUserResponse{Id: userID}, nil
}
