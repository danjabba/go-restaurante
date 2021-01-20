package routes

import (
	adminitrationapiv1 "github.com/danjabba/go-restaurante/internal/handlers/administration"
	echo "github.com/labstack/echo/v4"
)

// RegisterMenusRoutes ...
func RegisterMenusRoutes(e *echo.Echo) {

	e.GET("/api/v1/menus", adminitrationapiv1.ListMenu)
	e.POST("/api/v1/menus", adminitrationapiv1.CreateMenu)
	e.GET("/api/v1/menus/:identifier", adminitrationapiv1.GetMenu)
	e.PUT("/api/v1/menus/:identifier", adminitrationapiv1.UpdateMenu)
	e.DELETE("/api/v1/menus/:identifier", adminitrationapiv1.DeleteMenu)

}
