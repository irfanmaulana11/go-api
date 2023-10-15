package repositories

import (
	"context"
	"go-api/app/models"
)

type DBRepositoryInterface interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	UpdateUser(ctx context.Context, data models.User) error
}
