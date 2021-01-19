package routes

import (
	adminitrationapiv1 "github.com/danjabba/go-restaurante/internal/handlers/administration"
	echo "github.com/labstack/echo/v4"
)

// RegisterFoodsRoutes ...
func RegisterFoodsRoutes(e *echo.Echo) {

	e.GET("/api/v1/foods", adminitrationapiv1.ListFood)
	e.POST("/api/v1/foods", adminitrationapiv1.CreateFood)
	e.GET("/api/v1/foods/:identifier", adminitrationapiv1.GetFood)
	e.PUT("/api/v1/foods/:identifier", adminitrationapiv1.UpdateFood)
	e.DELETE("/api/v1/foods/:identifier", adminitrationapiv1.DeleteFood)

}
