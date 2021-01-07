package routes

import (
	adminitrationapiv1 "github.com/danjabba/go-restaurante/internal/handlers/administration"
	echo "github.com/labstack/echo/v4"
)

// RegisterUSersRoutes ...
func RegisterUSersRoutes(e *echo.Echo) {

	e.GET("/api/v1/users", adminitrationapiv1.ListUser)
	e.POST("/api/v1/users", adminitrationapiv1.CreateUser)
	e.GET("/api/v1/users/:identifier", adminitrationapiv1.GetUser)
	e.PUT("/api/v1/users/:identifier", adminitrationapiv1.UpdateUser)
	e.DELETE("/api/v1/users/:identifier", adminitrationapiv1.DeleteUser)

}
