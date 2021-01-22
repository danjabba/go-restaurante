package administration

import (
	"fmt"
	"net/http"

	models "github.com/danjabba/go-restaurante/internal/database/models"
	"github.com/danjabba/go-restaurante/tools"
	utils "github.com/danjabba/go-restaurante/tools"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

// createUserPayload ...
type createUserPayload struct {
	DniType   string `bson:"dni_type" json:"dni_type"`
	DniNumber string `bson:"dni_number" validate:"required,min=4,max=50" json:"dni_number"`
	Name      string `bson:"name" json:"name" validate:"required,min=4,max=50"`
	LastName  string `bson:"last_name" json:"last_name" validate:"required,min=4,max=50"`
	Email     string `bson:"email" json:"email" validate:"email"`
}

// updateUserPayload ...
type updateUserPayload struct {
	DniType   string `bson:"dni_type" json:"dni_type"`
	DniNumber string `bson:"dni_number" validate:"required,min=4,max=50" json:"dni_number"`
	Name      string `bson:"name" json:"name" validate:"required,min=4,max=50"`
	LastName  string `bson:"last_name" json:"last_name" validate:"required,min=4,max=50"`
	Email     string `bson:"email" json:"email" validate:"email"`
}

// ListUser ...
func ListUser(c echo.Context) error {

	db := tools.GetConnection()

	model := &models.User{}
	dataDB, err := model.Query(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/user.go: Error listing", err)
		return utils.ReturnError(err, c)
	}

	return c.JSON(http.StatusOK, utils.CreateResponse(dataDB, "Users are listed", true))

}

// CreateUser ...
func CreateUser(c echo.Context) error {

	db := tools.GetConnection()

	payload := new(createUserPayload)

	if err := c.Bind(payload); err != nil {

		fmt.Println("ERR handlers/administration/user.go: Creating - Binding payload", err)
		return c.JSON(http.StatusUnprocessableEntity, utils.CreateResponse(err, "internal error", false))

	}

	validate := validator.New()
	err := validate.Struct(payload)

	if err != nil {

		fmt.Println("ERR handlers/administration/user.go: Creating - Payload validation", err)
		validationerrors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(nil, validationerrors.Error(), false))

	}

	model := &models.User{
		DniType:   payload.DniType,
		DniNumber: payload.DniNumber,
		Name:      payload.Name,
		LastName:  payload.LastName,
		Email:     payload.Email,
	}

	err = model.Create(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/user.go: Creating - Saving model", err)
		return utils.ReturnError(err, c)
	}

	return c.JSON(http.StatusCreated, utils.CreateResponse(model, "User created", true))

}

// DeleteUser ...
func DeleteUser(c echo.Context) error {

	db := tools.GetConnection()

	identifier := c.Param("identifier")

	model := &models.User{}
	model.ID = identifier

	_, err := model.QueryByID(db)

	if err != nil {

		return c.JSON(http.StatusNotFound, utils.CreateResponse(nil, "User not found", true))
	}

	dataDB, err := model.Delete(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/user.go: Deleting - Deleting model", err)
		return utils.ReturnError(err, c)
	}

	return c.JSON(http.StatusOK, utils.CreateResponse(dataDB, "User deleted", true))

}

// GetUser ...
func GetUser(c echo.Context) error {

	db := tools.GetConnection()

	identifier := c.Param("identifier")

	model := &models.User{}
	model.ID = identifier

	dataDB, err := model.QueryByID(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/user.go: Getting - getting model: ", err)
		return c.JSON(http.StatusNotFound, utils.CreateResponse(nil, "User not found", true))
	}

	return c.JSON(http.StatusOK, utils.CreateResponse(dataDB, "User retrieved", true))

}

// UpdateUser ...
func UpdateUser(c echo.Context) error {

	db := tools.GetConnection()

	payload := new(updateUserPayload)

	if err := c.Bind(payload); err != nil {

		fmt.Println("ERR handlers/administration/user.go: Updating - Binding payload ", err)
		return c.JSON(http.StatusUnprocessableEntity, utils.CreateResponse(err, "internal error", false))

	}

	validate := validator.New()
	err := validate.Struct(payload)

	if err != nil {

		fmt.Println("ERR handlers/administration/user.go: Updating - Payload validation ", err)
		validationErrors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(nil, validationErrors.Error(), false))
	}

	identifier := c.Param("identifier")
	currentModel := &models.User{}
	currentModel.ID = identifier

	currentModel, err = currentModel.QueryByID(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/user.go: Updating - getting model", err)
		return c.JSON(http.StatusNotFound, utils.CreateResponse(nil, "User not found", true))

	}

	// New fields

	currentModel.Name = payload.Name
	currentModel.LastName = payload.LastName
	currentModel.DniType = payload.DniType
	currentModel.DniNumber = payload.DniNumber
	currentModel.Email = payload.Email

	dataDB, err := currentModel.Update(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/user.go: Updating - updating model", err)
		return utils.ReturnError(err, c)

	}

	return c.JSON(http.StatusOK, utils.CreateResponse(dataDB, "User Updated", true))
}
