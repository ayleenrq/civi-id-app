package main

import (
	"civi-id-app/configs"
	datasources "civi-id-app/internal/dataSources"
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
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH},
	}))

	for _, r := range e.Routes() {
		log.Printf("ROUTE %s %s", r.Method, r.Path)
	}

	cld, err := datasources.NewCloudinaryClient()
	if err != nil {
		log.Fatalf("Failed to init cloudinary client: %v", err)
	}

	routes.Routes(e, db, cld)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(e.Start(":" + port))
}
