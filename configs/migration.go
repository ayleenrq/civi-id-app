package configs

import (
	"dating-app/internal/models"
	"fmt"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	fmt.Println("Running migrations...")

	err := db.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		panic("Migration failed: " + err.Error())
	}

	fmt.Println("Migrations completed!")
}
