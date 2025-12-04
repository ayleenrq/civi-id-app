package adminservice

import (
	"civi-id-app/internal/models"
	adminrequest "civi-id-app/internal/dto/request/admin_request"
	"context"
)

type IAdminService interface {
	Register(ctx context.Context, req adminrequest.RegisterAdminRequest) error
	Login(ctx context.Context, req adminrequest.LoginAdminRequest) (string, error)
	GetProfile(ctx context.Context, adminId int) (*models.User, error)
	UpdateProfile(ctx context.Context, adminID int, req adminrequest.UpdateProfileRequest) error
	Logout(ctx context.Context, adminID int) error
}
