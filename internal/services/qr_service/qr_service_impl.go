package qrservice

import (
	"context"

	qrrequest "civi-id-app/internal/dto/request/qr_request"
	qrrepo "civi-id-app/internal/repositories/qr_repository"
	userrepo "civi-id-app/internal/repositories/user_repository"

	errorresponse "civi-id-app/pkg/constant/error_response"
)

type QRServiceImpl struct {
	qrRepo   qrrepo.IQRRepository
	userRepo userrepo.IUserRepository
}

func NewQRServiceImpl(qrRepo qrrepo.IQRRepository, userRepo userrepo.IUserRepository) IQRService {
	return &QRServiceImpl{qrRepo: qrRepo, userRepo: userRepo}
}

func (s *QRServiceImpl) Scan(ctx context.Context, req qrrequest.ScanQRRequest) (map[string]any, error) {
	session, err := s.qrRepo.FindByToken(ctx, req.QRToken)
	if err != nil {
		return nil, errorresponse.NewCustomError(errorresponse.ErrNotFound, "QR Code not found", 404)
	}

	user, err := s.userRepo.FindById(ctx, session.UserID)
	if err != nil {
		return nil, errorresponse.NewCustomError(errorresponse.ErrNotFound, "User not found", 404)
	}
	
	data := map[string]any{
		"id":           user.ID,
		"nik":          *user.NIK,
		"name":         user.Name,
		"email":        user.Email,
		"tempat_lahir": *user.TempatLahir,
		"birth_date":   user.BirthDate,
		"agama":        *user.Agama,
		"address":      *user.Address,
		"phone_number": *user.PhoneNumber,
		"status":       *user.Status,
		"gender_verified": func() bool {
			if user.GenderVerified == nil {
				return false
			}
			return *user.GenderVerified
		}(),
		"photo_url":  *user.PhotoURL,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	}

	return data, nil
}
