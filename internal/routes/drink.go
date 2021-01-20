package routes

import (
	adminitrationapiv1 "github.com/danjabba/go-restaurante/internal/handlers/administration"
	echo "github.com/labstack/echo/v4"
)

// RegisterDrinksRoutes ...
func RegisterDrinksRoutes(e *echo.Echo) {

	e.GET("/api/v1/drinks", adminitrationapiv1.ListDrink)
	e.POST("/api/v1/drinks", adminitrationapiv1.CreateDrink)
	e.GET("/api/v1/drinks/:identifier", adminitrationapiv1.GetDrink)
	e.PUT("/api/v1/drinks/:identifier", adminitrationapiv1.UpdateDrink)
	e.DELETE("/api/v1/drinks/:identifier", adminitrationapiv1.DeleteDrink)

}
