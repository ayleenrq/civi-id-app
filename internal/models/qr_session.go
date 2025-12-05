package models

import "time"

type QRSession struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int       `gorm:"index" json:"user_id"`
	QRToken   string    `gorm:"type:text" json:"qr_token"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
