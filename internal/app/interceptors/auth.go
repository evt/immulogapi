package interceptors

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Validator interface {
	Validate(token string) error
}

type Auth struct {
	validator Validator
}

func NewAuthInterceptor(validator Validator) *Auth {
	return &Auth{validator: validator}
}

func (a *Auth) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if info.FullMethod == "/v1.AuthService/Auth" {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "no metadata provided")
		}

		values := md["authorization"]
		if len(values) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
		}

		token := values[0]
		token = strings.TrimPrefix(token, "Bearer ")

		if err := a.validator.Validate(token); err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
		}

		return handler(ctx, req)
	}
}
