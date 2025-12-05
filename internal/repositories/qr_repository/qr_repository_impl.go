package qrrepository

import (
	"civi-id-app/internal/models"
	"context"

	"gorm.io/gorm"
)

type QRRepositoryImpl struct {
	db *gorm.DB
}

func NewQRRepositoryImpl(db *gorm.DB) IQRRepository {
	return &QRRepositoryImpl{db}
}

func (r *QRRepositoryImpl) Create(ctx context.Context, qr *models.QRSession) error {
	return r.db.WithContext(ctx).Create(qr).Error
}

func (r *QRRepositoryImpl) FindByToken(ctx context.Context, token string) (*models.QRSession, error) {
	var session models.QRSession
	err := r.db.WithContext(ctx).Where("qr_token = ?", token).First(&session).Error
	return &session, err
}
