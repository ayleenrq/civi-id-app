package adminservice

import (
	adminrequest "civi-id-app/internal/dto/request/admin_request"
	adminrepo "civi-id-app/internal/repositories/admin_repository"
	"civi-id-app/pkg/utils"
	"civi-id-app/internal/models"
	errorresponse "civi-id-app/pkg/constant/error_response"
	"context"
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type AdminServiceImpl struct {
	adminRepo adminrepo.IAdminRepository
}

func NewAdminServiceImpl(adminRepo adminrepo.IAdminRepository) IAdminService {
	return &AdminServiceImpl{adminRepo: adminRepo}
}

func (a *AdminServiceImpl) Register(ctx context.Context, req adminrequest.RegisterAdminRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Name is required", 400)
	}

	if strings.TrimSpace(req.Email) == "" {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Email is required", 400)
	}

	if strings.TrimSpace(req.Password) == "" {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Password is required", 400)
	}

	if !utils.IsValidEmail(req.Email) {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Email format is invalid", 400)
	}

	existsEmail, err := a.adminRepo.FindByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errorresponse.NewCustomError(errorresponse.ErrInternal, "Failed to check email", 500)
	}
	
	if existsEmail != nil {
		return errorresponse.NewCustomError(errorresponse.ErrExists, "Email already exists", 409)
	}

	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Failed to hash password", 400)
	}

	role, err := a.adminRepo.FindRoleAdmin(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errorresponse.NewCustomError(errorresponse.ErrNotFound, "Role 'admin' not found. Please create it first in /api/v1/role/create", 404)
		}
		return errorresponse.NewCustomError(errorresponse.ErrInternal, "Failed to get role admin", 500)
	}

	admin := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashed,
		RoleID:   int(role.ID),
		Role:     *role,
	}

	if err := a.adminRepo.Create(ctx, admin); err != nil {
		return errorresponse.NewCustomError(errorresponse.ErrInternal, "Failed to create admin", 500)
	}

	return nil
}

func (a *AdminServiceImpl) Login(ctx context.Context, req adminrequest.LoginAdminRequest) (string, error) {
	admin, err := a.adminRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return "", errorresponse.NewCustomError(errorresponse.ErrNotFound, "Invalid credentials", 400)
	}

	isPasswordValid := utils.CheckPasswordHash(req.Password, admin.Password)
	if !isPasswordValid {
		return "", errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Password incorrect", 400)
	}

	token, err := utils.GenerateToken(admin.ID, admin.RoleID)
	if err != nil {
		return "", errorresponse.NewCustomError(errorresponse.ErrInternal, "Failed to generate token", 500)
	}

	return token, nil
}

func (a *AdminServiceImpl) GetProfile(ctx context.Context, adminId int) (*models.User, error) {
	admin, err := a.adminRepo.FindByAdminID(ctx, adminId)
	if err != nil {
		return nil, errorresponse.NewCustomError(errorresponse.ErrNotFound, "Admin not found", 404)
	}
	return admin, nil
}

func (a *AdminServiceImpl) UpdateProfile(ctx context.Context, adminID int, req adminrequest.UpdateProfileRequest) error {
	admin, err := a.adminRepo.FindByAdminID(ctx, adminID)
	if err != nil {
		return errorresponse.NewCustomError(errorresponse.ErrNotFound, "Admin not found", 404)
	}

	if strings.TrimSpace(req.Name) == "" {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Name is required", 400)
	}

	admin.Name = req.Name

	if err := a.adminRepo.Update(ctx, admin); err != nil {
		return errorresponse.NewCustomError(errorresponse.ErrInternal, "Failed to update admin profile", 500)
	}

	return nil
}

func (a *AdminServiceImpl) Logout(ctx context.Context, adminID int) error {
	fmt.Printf("Admin %d logged out\n", adminID)
	return nil
}
