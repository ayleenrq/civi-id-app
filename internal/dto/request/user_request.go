package request

import "time"

type RegisterUserRequest struct {
	NIK          string    `json:"nik" form:"nik"`
	Name         string    `json:"name" form:"name"`
	Email        string    `json:"email" form:"email"`
	Password     string    `json:"password" form:"password"`
	JenisKelamin string    `json:"jenis_kelamin" form:"jenis_kelamin"`
	TempatLahir  string    `json:"tempat_lahir" form:"tempat_lahir"`
	BirthDate    time.Time `json:"birth_date" form:"birth_date"`
	Agama        string    `json:"agama" form:"agama"`
	Address      string    `json:"address" form:"address"`
	PhoneNumber  string    `json:"phone_number" form:"phone_number"`
	PhotoFile    string    `json:"photo_file" form:"photo"`
}

type LoginUserRequest struct {
	NIK      string `json:"nik" form:"nik"`
	Password string `json:"password" form:"password"`
}
