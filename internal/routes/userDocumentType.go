package routes

import (
	adminitrationapiv1 "github.com/danjabba/go-restaurante/internal/handlers/administration"
	echo "github.com/labstack/echo/v4"
)

// RegisterUserDocumentTypesRoutes ...
func RegisterUserDocumentTypesRoutes(e *echo.Echo) {

	e.GET("/api/v1/userdocumenttypes", adminitrationapiv1.ListUserDocumentType)
	e.POST("/api/v1/userdocumenttypes", adminitrationapiv1.CreateUserDocumentType)
	e.GET("/api/v1/userdocumenttypes/:identifier", adminitrationapiv1.GetUserDocumentType)
	e.PUT("/api/v1/userdocumenttypes/:identifier", adminitrationapiv1.UpdateUserDocumentType)
	e.DELETE("/api/v1/userdocumenttypes/:identifier", adminitrationapiv1.DeleteUserDocumentType)

}
