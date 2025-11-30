package configs

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func InitDB() *gorm.DB {
	once.Do(func() {
		dsn := os.Getenv("DATABASE_URL")
		if dsn == "" {
			log.Fatal("DATABASE_URL is required but not set")
		}

		var err error
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to MySQL: %v", err)
		}

		fmt.Println("MySQL connected successfully!")
	})

	return DB
}

func CloseConnections() {
	if DB == nil {
		return
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Println("Failed to get SQL DB:", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Println("Failed to close DB:", err)
	} else {
		fmt.Println("Database connection closed")
	}
}
