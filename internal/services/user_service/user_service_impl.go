package userservice

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	userrequest "civi-id-app/internal/dto/request/user_request"
	"civi-id-app/internal/models"
	qrrepo "civi-id-app/internal/repositories/qr_repository"
	userrepo "civi-id-app/internal/repositories/user_repository"
	errorresponse "civi-id-app/pkg/constant/error_response"
	"civi-id-app/pkg/utils"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	userRepo   userrepo.IUserRepository
	qrRepo     qrrepo.IQRRepository
	cloudinary *cloudinary.Cloudinary
}

func NewUserServiceImpl(userRepo userrepo.IUserRepository, qrRepo qrrepo.IQRRepository, cld *cloudinary.Cloudinary) IUserService {
	return &UserServiceImpl{userRepo: userRepo, qrRepo: qrRepo, cloudinary: cld}
}

func (s *UserServiceImpl) Register(ctx context.Context, req userrequest.RegisterUserRequest) error {
	if strings.TrimSpace(req.NIK) == "" {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "NIK is required", 400)
	}
	if strings.TrimSpace(req.Name) == "" {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Name is required", 400)
	}
	if strings.TrimSpace(req.Email) == "" {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Email is required", 400)
	}
	if strings.TrimSpace(req.Password) == "" {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Password is required", 400)
	}

	if strings.TrimSpace(req.TempatLahir) == "" {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Tempat lahir is required", 400)
	}
	if strings.TrimSpace(req.Agama) == "" {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Agama is required", 400)
	}
	if strings.TrimSpace(req.Address) == "" {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Address is required", 400)
	}
	if strings.TrimSpace(req.PhoneNumber) == "" {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Phone number is required", 400)
	}
	if strings.TrimSpace(req.Status) == "" {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Status is required", 400)
	}

	if !utils.IsValidEmail(req.Email) {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Email format is invalid", 400)
	}

	if !utils.IsValidNIK(req.NIK) {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "NIK must be exactly 16 digits", 400)
	}

	if !utils.IsNumeric(req.PhoneNumber) {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Phone number must be numeric", 400)
	}

	if req.PhotoFile == nil {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Photo file is required", 400)
	}

	existsNIK, err := s.userRepo.FindByNIK(ctx, req.NIK)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errorresponse.NewCustomError(errorresponse.ErrInternal, "Failed to get user NIK", 500)
	}
	if existsNIK != nil {
		return errorresponse.NewCustomError(errorresponse.ErrExists, "NIK already exists", 409)
	}

	hashedPass, err := utils.HashPassword(req.Password)
	if err != nil {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Failed to hash password", 400)
	}

	role, err := s.userRepo.FindRoleUser(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errorresponse.NewCustomError(errorresponse.ErrNotFound, "Role 'user' not found.", 404)
		}
		return errorresponse.NewCustomError(errorresponse.ErrInternal, "Failed to get role user", 500)
	}

	birth, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "BirthDate must be YYYY-MM-DD", 400)
	}

	photoURL, err := utils.UploadToCloudinary(req.PhotoFile, "civi-id/users")
	if err != nil {
		return errorresponse.NewCustomError(errorresponse.ErrInternal, "Failed to upload photo", 500)
	}

	genderML, err := utils.DetectGenderML(req.PhotoFile)
	if err != nil {
		return errorresponse.NewCustomError(errorresponse.ErrInternal, "Failed to detect gender", 500)
	}

	jenisKelamin := utils.MLToIndo(genderML)

	user := &models.User{
		NIK:            &req.NIK,
		Name:           req.Name,
		Email:          req.Email,
		Password:       hashedPass,
		JenisKelamin:   &jenisKelamin,
		TempatLahir:    &req.TempatLahir,
		BirthDate:      &birth,
		Agama:          &req.Agama,
		Address:        &req.Address,
		PhoneNumber:    &req.PhoneNumber,
		Status:         &req.Status,
		PhotoURL:       &photoURL,
		GenderML:       &genderML,
		RoleID:         role.ID,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return errorresponse.NewCustomError(errorresponse.ErrInternal, "Failed to create user", 500)
	}

	return nil
}

func (s *UserServiceImpl) Login(ctx context.Context, req userrequest.LoginUserRequest) (string, error) {
	user, err := s.userRepo.FindByNIK(ctx, req.NIK)
	if err != nil {
		return "", errorresponse.NewCustomError(errorresponse.ErrNotFound, "Invalid NIK", 400)
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return "", errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Password incorrect", 400)
	}

	token, err := utils.GenerateToken(user.ID, user.RoleID)
	if err != nil {
		return "", errorresponse.NewCustomError(errorresponse.ErrInternal, "Failed to generate token", 500)
	}

	return token, nil
}

func (s *UserServiceImpl) GetProfile(ctx context.Context, userID int) (*models.User, error) {
	user, err := s.userRepo.FindById(ctx, userID)
	if err != nil {
		return nil, errorresponse.NewCustomError(errorresponse.ErrNotFound, "User not found", 404)
	}
	return user, nil
}

func (s *UserServiceImpl) UpdateProfile(ctx context.Context, userID int, req userrequest.UpdateUserRequest) error {
	user, err := s.userRepo.FindById(ctx, userID)
	if err != nil {
		return errorresponse.NewCustomError(errorresponse.ErrNotFound, "User not found", 404)
	}

	if req.Email != "" {
		user.Email = req.Email
	}

	if req.Address != "" {
		user.Address = &req.Address
	}

	if req.PhoneNumber != "" {
		user.PhoneNumber = &req.PhoneNumber
	}

	if err := s.userRepo.Update(ctx, user); err != nil {
		return errorresponse.NewCustomError(errorresponse.ErrInternal, "Failed to update user", 500)
	}

	return nil
}

func (s *UserServiceImpl) GenerateQR(ctx context.Context, userID int) (string, error) {
    user, err := s.userRepo.FindById(ctx, userID)
    if err != nil {
        return "", errorresponse.NewCustomError(errorresponse.ErrNotFound, "User not found", 404)
    }

    // Generate token baru setiap user klik Generate QR
    qrToken := uuid.NewString()

    // Buat QR code (bytes)
    qrBytes, err := utils.GenerateQRCodeBytes(qrToken)
    if err != nil {
        return "", errorresponse.NewCustomError(errorresponse.ErrInternal, "Failed to generate QR code", 500)
    }

    // Upload ke Cloudinary
    filename := fmt.Sprintf("qr_user_%d_%s", user.ID, qrToken)

    uploadResp, err := s.cloudinary.Upload.Upload(
        ctx,
        bytes.NewReader(qrBytes),
        uploader.UploadParams{
            PublicID:     "civi-id/qr/" + filename,
            Folder:       "civi-id/qr",
            ResourceType: "image",
            Overwrite:    boolPtr(true),
        },
    )
    if err != nil {
        return "", errorresponse.NewCustomError(errorresponse.ErrInternal, "Failed to upload QR code", 500)
    }

    // Simpan token baru ke database
    session := models.QRSession{
        UserID:  user.ID,
        QRToken: qrToken,
    }
    if err := s.qrRepo.Create(ctx, &session); err != nil {
        return "", errorresponse.NewCustomError(errorresponse.ErrInternal, "Failed to save QR session", 500)
    }

    return uploadResp.SecureURL, nil
}


func (s *UserServiceImpl) Logout(ctx context.Context, userID int) error {
	fmt.Printf("User %d logged out\n", userID)
	return nil
}

func boolPtr(b bool) *bool {
	return &b
}
