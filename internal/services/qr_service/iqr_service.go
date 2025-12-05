package qrservice

import (
	qrrequest "civi-id-app/internal/dto/request/qr_request"
	"context"
)

type IQRService interface {
	Scan(ctx context.Context, req qrrequest.ScanQRRequest) (map[string]any, error)
}
