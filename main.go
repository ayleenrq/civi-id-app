package main

import (
	"civi-id-app/configs"
	"civi-id-app/routes"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	configs.LoadEnv()

	db := configs.InitDB()

	configs.RunMigrations(db)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH,
		},
	}))

	routes.Routes(e, db)

	for _, r := range e.Routes() {
		log.Printf("ROUTE %s %s", r.Method, r.Path)
	}

	defer configs.CloseConnections()

	// Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Server running on port:", port)

	log.Fatal(e.Start(":" + port))
}
