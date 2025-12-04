package userrepository

import (
	"civi-id-app/internal/models"
	"context"
)

type IUserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByNIK(ctx context.Context, nik string) (*models.User, error)
	FindRoleUser(ctx context.Context) (*models.Role, error)
	FindById(ctx context.Context, userID int) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
}
