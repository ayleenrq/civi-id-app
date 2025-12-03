package rolerepository

import (
	"civi-id-app/internal/models"
	"context"
)

type IRoleRepository interface {
	Create(ctx context.Context, data *models.Role) error
	FindByName(ctx context.Context, name string) (*models.Role, error)
	FindAll(ctx context.Context, limit, offset int, search string) ([]*models.Role, int, error)
	FindById(ctx context.Context, roleId int) (*models.Role, error)
	Update(ctx context.Context, roleId int, data *models.Role) error
	Delete(ctx context.Context, roleId int) error
}
