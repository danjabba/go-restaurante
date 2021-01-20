package main

import (
	"os"

	routing "github.com/danjabba/go-restaurante/internal/routes"
	"github.com/danjabba/go-restaurante/tools"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	_ = tools.GetConnection()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*", "http://localhost:3000", "http://127.0.0.1:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			echo.HeaderOrigin,
			echo.HeaderAllow,
			echo.HeaderAccessControlAllowHeaders,
		},
	}))

	routing.LoadRoutes(e)

	port := os.Getenv("APP_PORT")

	e.Logger.Fatal(e.Start(":" + port))

}
