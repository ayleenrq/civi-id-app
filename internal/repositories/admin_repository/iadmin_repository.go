package adminrepository

import (
	"civi-id-app/internal/models"
	"context"
)

type IAdminRepository interface {
	Create(ctx context.Context, admin *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindRoleAdmin(ctx context.Context) (*models.Role, error)
	FindByAdminID(ctx context.Context, adminID int) (*models.User, error)
	Update(ctx context.Context, admin *models.User) error
}
