package adminrepository

import (
	"civi-id-app/internal/models"
	"context"

	"gorm.io/gorm"
)

type AdminRepositoryImpl struct {
	DB *gorm.DB
}

func NewAdminRepositoryImpl(db *gorm.DB) IAdminRepository {
	return &AdminRepositoryImpl{DB: db}
}

func (r *AdminRepositoryImpl) Create(ctx context.Context, admin *models.User) error {
	return r.DB.WithContext(ctx).Create(admin).Error
}

func (r *AdminRepositoryImpl) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var admin models.User
	if err := r.DB.WithContext(ctx).First(&admin, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (a *AdminRepositoryImpl) FindRoleAdmin(ctx context.Context) (*models.Role, error) {
	var role models.Role
	if err := a.DB.WithContext(ctx).Where("LOWER(name) = LOWER(?)", "admin").First(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *AdminRepositoryImpl) FindByAdminID(ctx context.Context, adminID int) (*models.User, error) {
	var admin models.User
	if err := r.DB.WithContext(ctx).Preload("Role").First(&admin, "id = ?", adminID).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *AdminRepositoryImpl) Update(ctx context.Context, admin *models.User) error {
	return r.DB.WithContext(ctx).Save(admin).Error
}
