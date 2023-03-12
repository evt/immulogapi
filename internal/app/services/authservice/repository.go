package authservice

import (
	"context"

	"github.com/evt/immulogapi/internal/app/models"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.User) (int64, error)
	GetUser(ctx context.Context, user, pass string) (*models.User, error)
}
