package qrrepository

import (
	"civi-id-app/internal/models"
	"context"
)

type IQRRepository interface {
	Create(ctx context.Context, qr *models.QRSession) error
	FindByToken(ctx context.Context, token string) (*models.QRSession, error)
}