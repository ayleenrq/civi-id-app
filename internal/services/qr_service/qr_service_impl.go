package qrservice

import (
	"context"

	qrrequest "civi-id-app/internal/dto/request/qr_request"
	scanqrresponse "civi-id-app/internal/dto/response/scanqr_response"
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

func (s *QRServiceImpl) Scan(ctx context.Context, req qrrequest.ScanQRRequest) (*scanqrresponse.ScanQRResponse, error) {
	session, err := s.qrRepo.FindByToken(ctx, req.QRToken)
	if err != nil {
		return nil, errorresponse.NewCustomError(errorresponse.ErrNotFound, "QR Code not found", 404)
	}

	user, err := s.userRepo.FindById(ctx, session.UserID)
	if err != nil {
		return nil, errorresponse.NewCustomError(errorresponse.ErrNotFound, "User not found", 404)
	}

	response := scanqrresponse.ToScanQRResponse(*user)

	return &response, nil
}

