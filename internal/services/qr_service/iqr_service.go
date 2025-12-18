package qrservice

import (
	qrrequest "civi-id-app/internal/dto/request/qr_request"
	scanqrresponse "civi-id-app/internal/dto/response/scanqr_response"
	"context"
)

type IQRService interface {
	Scan(ctx context.Context, req qrrequest.ScanQRRequest) (*scanqrresponse.ScanQRResponse, error)
}
