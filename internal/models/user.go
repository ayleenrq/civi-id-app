package models

import "time"

type User struct {
	ID             int        `gorm:"primaryKey;autoIncrement" json:"id"`
	NIK            *string    `gorm:"type:varchar(255)" json:"nik"`
	Name           string     `gorm:"type:varchar(255)" json:"name"`
	Email          string     `gorm:"type:varchar(255)" json:"email"`
	Password       string     `gorm:"type:varchar(255)" json:"password"`
	JenisKelamin   *string    `gorm:"type:varchar(50)" json:"jenis_kelamin"`
	TempatLahir    *string    `gorm:"type:varchar(255)" json:"tempat_lahir"`
	BirthDate      *time.Time `gorm:"type:date" json:"birth_date"`
	Agama          *string    `gorm:"type:varchar(100)" json:"agama"`
	Address        *string    `gorm:"type:text" json:"address"`
	PhoneNumber    *string    `gorm:"type:varchar(20)" json:"phone_number"`
	Status         *string    `gorm:"type:varchar(50)" json:"status"`
	GenderVerified *bool      `gorm:"type:boolean" json:"gender_verified"`
	GenderML       *string    `gorm:"type:varchar(20)" json:"gender_ml"`
	RoleID         int        `gorm:"type:int" json:"role_id"`
	Role           Role       `gorm:"foreignKey:RoleID"`
	PhotoURL       *string    `gorm:"type:text" json:"photo_url"`
	CreatedAt      time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}
