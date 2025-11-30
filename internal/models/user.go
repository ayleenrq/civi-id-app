package models

import "time"

type User struct {
	ID             string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(255)" json:"name"`
	Email          string    `gorm:"type:varchar(255);uniqueIndex" json:"email"`
	Password       string    `gorm:"type:varchar(255)" json:"password"`
	Age            int       `gorm:"type:int" json:"age"`
	GenderIdentity string    `gorm:"type:varchar(20)" json:"gender_identity"`
	GenderVerified bool      `gorm:"type:boolean" json:"gender_verified"`
	GenderML       string    `gorm:"type:varchar(20)" json:"gender_ml"`
	Bio            string    `gorm:"type:text" json:"bio"`
	Hobby          string    `gorm:"type:varchar(255)" json:"hobby"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}