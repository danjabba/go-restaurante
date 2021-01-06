package tools

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"

	gorm "github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

var (
	// ErrConflict Conflict
	ErrConflict = errors.New("conflict")
	// ErrAlredyExists Already exists
	ErrAlredyExists = errors.New("already exists")
)

// ReturnError ...
func ReturnError(err error, c echo.Context) error {
	code, body := GetHTTPError(err)
	return c.JSON(code, body)
}

// GetHTTPError ...
func GetHTTPError(err error) (int, ResponseStructure) {

	fmt.Println("...internal error ", err)

	switch reflect.TypeOf(err) {

	case reflect.TypeOf(gorm.Errors{}):
		{

			if errors.Is(err, gorm.ErrRecordNotFound) {
				return http.StatusNotFound, CreateResponse(nil, err.Error(), false)
			}

		}

	}

	return http.StatusInternalServerError, CreateResponse(nil, "Internal error", false)

}
