package routes

import (
	adminitrationapiv1 "github.com/danjabba/go-restaurante/internal/handlers/administration"
	echo "github.com/labstack/echo/v4"
)

// RegisterRestaurantsRoutes ...
func RegisterRestaurantsRoutes(e *echo.Echo) {

	e.GET("/api/v1/restaurants", adminitrationapiv1.ListRestaurant)
	e.POST("/api/v1/restaurants", adminitrationapiv1.CreateRestaurant)
	e.GET("/api/v1/restaurants/:identifier", adminitrationapiv1.GetRestaurant)
	e.PUT("/api/v1/restaurants/:identifier", adminitrationapiv1.UpdateRestaurant)
	e.DELETE("/api/v1/restaurants/:identifier", adminitrationapiv1.DeleteRestaurant)

}
