package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB) {
	// Define your routes here, for example:
	// e.GET("/users", handlers.GetUsers(db))
}
