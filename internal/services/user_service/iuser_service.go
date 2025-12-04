package userservice

import (
	"civi-id-app/internal/models"
	userrequest "civi-id-app/internal/dto/request/user_request"
	"context"
)

type IUserService interface {
	Register(ctx context.Context, req userrequest.RegisterUserRequest) error
	Login(ctx context.Context, req userrequest.LoginUserRequest) (string, error)
	GetProfile(ctx context.Context, userID int) (*models.User, error)
	UpdateProfile(ctx context.Context, adminID int, req userrequest.UpdateUserRequest) error
	Logout(ctx context.Context, userID int) error
}
