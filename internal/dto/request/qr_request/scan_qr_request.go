package qrrequest

type ScanQRRequest struct {
	QRToken string `json:"qr_token" form:"qr_token"`
}
