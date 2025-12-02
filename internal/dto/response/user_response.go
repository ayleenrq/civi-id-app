package response

import (
	"civi-id-app/internal/models"
	"civi-id-app/pkg/utils"
	"time"
)

type UserResponse struct {
	ID             int       `json:"id"`
	NIK            string    `json:"nik"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	TempatLahir    string    `json:"tempat_lahir"`
	BirthDate      time.Time `json:"birth_date"`
	Agama          string    `json:"agama"`
	Address        string    `json:"address"`
	PhoneNumber    string    `json:"phone_number"`
	PhotoURL       string    `json:"photo_url"`
	GenderVerified bool      `json:"gender_verified"`
	Role           int       `json:"role"`
	CreatedAt      string    `json:"created_at"`
	UpdatedAt      string    `json:"updated_at"`
}

func ToUserResponse(user models.User) UserResponse {
	return UserResponse{
		ID:             user.ID,
		NIK:            user.NIK,
		Name:           user.Name,
		Email:          user.Email,
		TempatLahir:    user.TempatLahir,
		BirthDate:      user.BirthDate,
		Agama:          user.Agama,
		Address:        user.Address,
		PhoneNumber:    user.PhoneNumber,
		PhotoURL:       user.PhotoURL,
		GenderVerified: user.GenderVerified,
		Role:           user.RoleID,
		CreatedAt:      utils.FormatDate(user.CreatedAt),
		UpdatedAt:      utils.FormatDate(user.UpdatedAt),
	}
}
