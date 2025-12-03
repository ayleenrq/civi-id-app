package rolerepository

import (
	"civi-id-app/internal/models"
	"context"

	"gorm.io/gorm"
)

type RoleRepositoryImpl struct {
	DB *gorm.DB
}

func NewRoleRepositoryImpl(db *gorm.DB) *RoleRepositoryImpl {
	return &RoleRepositoryImpl{DB: db}
}

func (r *RoleRepositoryImpl) Create(ctx context.Context, data *models.Role) error {
	return r.DB.WithContext(ctx).Create(data).Error
}

func (r *RoleRepositoryImpl) FindByName(ctx context.Context, name string) (*models.Role, error) {
	var role models.Role

	if err := r.DB.WithContext(ctx).Where("LOWER(name) = LOWER(?)", name).First(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *RoleRepositoryImpl) FindAll(ctx context.Context, limit, offset int, search string) ([]*models.Role, int, error) {
	var (
		roles []*models.Role
		count int64
	)

	query := r.DB.WithContext(ctx).Model(&models.Role{})

	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(limit).Offset(offset).Order("created_at DESC").Find(&roles).Error; err != nil {
		return nil, 0, err
	}

	return roles, int(count), nil
}

func (r *RoleRepositoryImpl) FindById(ctx context.Context, roleId int) (*models.Role, error) {
	var role models.Role

	if err := r.DB.WithContext(ctx).Where("id = ?", roleId).First(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *RoleRepositoryImpl) Update(ctx context.Context, roleId int, data *models.Role) error {
	var existing models.Role

	if err := r.DB.WithContext(ctx).Where("id = ?", roleId).First(&existing).Error; err != nil {
		return err
	}

	existing.Name = data.Name
	return r.DB.WithContext(ctx).Save(&existing).Error
}

func (r *RoleRepositoryImpl) Delete(ctx context.Context, roleId int) error {
	return r.DB.WithContext(ctx).Delete(&models.Role{}, "id = ?", roleId).Error
}
