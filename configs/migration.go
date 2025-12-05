package configs

import (
	"civi-id-app/internal/models"
	"fmt"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	fmt.Println("Running migrations...")

	err := db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.QRSession{},
	)

	if err != nil {
		panic("Migration failed: " + err.Error())
	}

	fmt.Println("Migrations completed!")
}
