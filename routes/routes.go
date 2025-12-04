package routes

import (
	"civi-id-app/routes/admin_route"
	"civi-id-app/routes/role_route"
	"civi-id-app/routes/user_route"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB) {
	v1 := e.Group("/api/v1")

	role_route.RoleRoutes(v1.Group("/role"), db)
	admin_route.AdminRoutes(v1.Group("/admin"), db)
	user_route.UserRoutes(v1.Group("/user"), db)
}
