package userrepository

import (
	"civi-id-app/internal/models"
	"context"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) IUserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, user *models.User) error {
	return r.DB.WithContext(ctx).Create(user).Error
}

func (r *UserRepositoryImpl) FindByNIK(ctx context.Context, nik string) (*models.User, error) {
	var user models.User
	if err := r.DB.WithContext(ctx).Where("nik = ?", nik).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) FindRoleUser(ctx context.Context) (*models.Role, error) {
	var role models.Role
	if err := r.DB.WithContext(ctx).Where("LOWER(name) = LOWER(?)", "user").First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *UserRepositoryImpl) FindById(ctx context.Context, userID int) (*models.User, error) {
	var user models.User
	if err := r.DB.WithContext(ctx).Preload("Role").First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, user *models.User) error {
	return r.DB.WithContext(ctx).Save(user).Error
}
