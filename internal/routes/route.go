package routes

import (
	echo "github.com/labstack/echo/v4"
)

// LoadRoutes ...
func LoadRoutes(e *echo.Echo) {

	RegisterUserDocumentTypesRoutes(e)
	RegisterUSersRoutes(e)
}
