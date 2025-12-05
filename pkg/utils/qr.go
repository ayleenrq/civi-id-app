package utils

import (
	"github.com/skip2/go-qrcode"
)

func GenerateQRCodeBytes(content string) ([]byte, error) {
	return qrcode.Encode(content, qrcode.Medium, 256)
}
