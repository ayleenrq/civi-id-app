package scanqrresponse

import (
	"civi-id-app/internal/models"
	"time"
)

type ScanQRResponse struct {
	ID             int       `json:"id"`
	NIK            string    `json:"nik"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	JenisKelamin   string    `json:"jenis_kelamin"`
	TempatLahir    string    `json:"tempat_lahir"`
	BirthDate      time.Time `json:"birth_date"`
	Agama          string    `json:"agama"`
	Address        string    `json:"address"`
	PhoneNumber    string    `json:"phone_number"`
	Status         string    `json:"status"`
	ReasonRegister string    `json:"alasan_register"`
	PhotoURL       string    `json:"photo_url"`
}

func ToScanQRResponse(user models.User) ScanQRResponse {
	return ScanQRResponse{
		ID:             user.ID,
		NIK:            *user.NIK,
		Name:           user.Name,
		Email:          user.Email,
		JenisKelamin:   *user.JenisKelamin,
		TempatLahir:    *user.TempatLahir,
		BirthDate:      *user.BirthDate,
		Agama:          *user.Agama,
		Address:        *user.Address,
		PhoneNumber:    *user.PhoneNumber,
		Status:         *user.Status,
		ReasonRegister: *user.ReasonRegister,
		PhotoURL:       *user.PhotoURL,
	}
}
