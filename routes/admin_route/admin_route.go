package admin_route

import (
	adminhandler "civi-id-app/internal/handlers/admin_handler"
	adminrepository "civi-id-app/internal/repositories/admin_repository"
	adminservice "civi-id-app/internal/services/admin_service"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AdminRoutes(e *echo.Group, db *gorm.DB) {
	adminRepo := adminrepository.NewAdminRepositoryImpl(db)
	adminService := adminservice.NewAdminServiceImpl(adminRepo)
	adminHandler := adminhandler.NewAdminHandler(adminService)

	e.POST("/register", adminHandler.RegisterAdmin)
	e.POST("/login", adminHandler.LoginAdmin)

	auth := e.Group("", echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	auth.GET("/me", adminHandler.GetProfileAdmin)
	auth.PUT("/me", adminHandler.UpdateProfileAdmin)
	auth.POST("/logout", adminHandler.LogoutAdmin)
}
