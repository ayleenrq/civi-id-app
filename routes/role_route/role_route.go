package role_route

import (
	rolehandler "civi-id-app/internal/handlers/role_handler"
	rolerepository "civi-id-app/internal/repositories/role_repository"
	roleservice "civi-id-app/internal/services/role_service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RoleRoutes(e *echo.Group, db *gorm.DB) {
	roleRepo := rolerepository.NewRoleRepositoryImpl(db)
	roleService := roleservice.NewRoleServiceImpl(roleRepo)
	roleHandler := rolehandler.NewRoleHandler(roleService)

	e.POST("/create", roleHandler.CreateRole)
    e.GET("/all", roleHandler.GetAllRole)
    e.GET("/:roleId", roleHandler.GetByIdRole)
    e.PUT("/:roleId/edit", roleHandler.UpdateRole)
    e.DELETE("/:roleId/delete", roleHandler.DeleteRole)
}
