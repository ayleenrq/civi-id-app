package user_route

import (
	userhandler "civi-id-app/internal/handlers/user_handler"
	qrrepository "civi-id-app/internal/repositories/qr_repository"
	userrepository "civi-id-app/internal/repositories/user_repository"
	userservice "civi-id-app/internal/services/user_service"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRoutes(e *echo.Group, db *gorm.DB, cld *cloudinary.Cloudinary) {
	userRepo := userrepository.NewUserRepositoryImpl(db)
	qrRepo := qrrepository.NewQRRepositoryImpl(db)
	userService := userservice.NewUserServiceImpl(userRepo, qrRepo, cld)
	userHandler := userhandler.NewUserHandler(userService)

	e.POST("/register", userHandler.RegisterUser)
	e.POST("/login", userHandler.LoginUser)

	auth := e.Group("", echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	auth.GET("/me", userHandler.GetProfileUser)
	auth.PUT("/me", userHandler.UpdateProfileUser)
	auth.GET("/qr", userHandler.GenerateQR)
	auth.POST("/logout", userHandler.LogoutUser)
}
