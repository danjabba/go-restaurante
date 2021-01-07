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

// createUserDocumentTypePayload ...
type createUserDocumentTypePayload struct {
	Label        string `bson:"label" json:"label" validate:"required,min=4,max=50"`
	Abbreviation string `bson:"abbreviation" json:"abbreviation" validate:"required,min=1,max=50"`
}

// updateUserDocumentTypePayload ...
type updateUserDocumentTypePayload struct {
	Label        string `bson:"label" json:"label" validate:"required,min=4,max=50"`
	Abbreviation string `bson:"abbreviation" json:"abbreviation" validate:"required,min=1,max=50"`
}

// ListUserDocumentType ...
func ListUserDocumentType(c echo.Context) error {

	db := tools.GetConnection()

	model := &models.UserDocumentType{}
	dataDB, err := model.Query(db)
	if err != nil {
		fmt.Println("ERR handlers/administration/userDocumentType.go: Error listing ", err)
		return utils.ReturnError(err, c)
	}

	return c.JSON(http.StatusOK, utils.CreateResponse(dataDB, "User Document types are listed", true))

}

// CreateUserDocumentType ...
func CreateUserDocumentType(c echo.Context) error {

	db := tools.GetConnection()

	payload := new(createUserDocumentTypePayload)

	if err := c.Bind(payload); err != nil {
		fmt.Println("ERR handlers/administration/userDocumentType.go: Creating - Binding payload ", err)
		return c.JSON(http.StatusUnprocessableEntity, utils.CreateResponse(err, "Internal error", false))
	}

	validate := validator.New()
	err := validate.Struct(payload)

	if err != nil {
		fmt.Println("ERR handlers/administration/userDocumentType.go: Creating - Payload validation ", err)
		validationErrors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(nil, validationErrors.Error(), false))
	}

	model := &models.UserDocumentType{
		Label:        payload.Label,
		Abbreviation: payload.Abbreviation,
	}

	err = model.Create(db)
	if err != nil {
		fmt.Println("ERR handlers/administration/condominiumType.go: Creating - Saving model ", err)
		return utils.ReturnError(err, c)
	}

	return c.JSON(http.StatusCreated, utils.CreateResponse(model, "User document type created", true))

}

// DeleteUserDocumentType ...
func DeleteUserDocumentType(c echo.Context) error {

	db := tools.GetConnection()

	identifier := c.Param("identifier")

	model := &models.UserDocumentType{}
	model.ID = identifier

	_, err := model.QueryByID(db)

	if err != nil {

		return c.JSON(http.StatusNotFound, utils.CreateResponse(nil, "User Document Type not found", true))
	}

	dataDB, err := model.Delete(db)
	if err != nil {
		fmt.Println("ERR handlers/administration/userDocumentType.go: Deleting - Deleting model ", err)
		return utils.ReturnError(err, c)
	}

	return c.JSON(http.StatusOK, utils.CreateResponse(dataDB, "User Document type deleted", true))

}

// GetUserDocumentType ...
func GetUserDocumentType(c echo.Context) error {

	db := tools.GetConnection()

	identifier := c.Param("identifier")

	model := &models.UserDocumentType{}
	model.ID = identifier

	dataDB, err := model.QueryByID(db)
	if err != nil {
		fmt.Println("ERR handlers/administration/userDocumentType.go: Getting - getting model ", err)
		return c.JSON(http.StatusNotFound, utils.CreateResponse(nil, "User Document Type not found", true))
	}

	return c.JSON(http.StatusOK, utils.CreateResponse(dataDB, "User Document type retrieved", true))

}

// UpdateUserDocumentType ...
func UpdateUserDocumentType(c echo.Context) error {

	db := tools.GetConnection()

	payload := new(updateUserDocumentTypePayload)

	if err := c.Bind(payload); err != nil {
		fmt.Println("ERR handlers/administration/userDocumentType.go: Updating - Binding payload ", err)
		return c.JSON(http.StatusUnprocessableEntity, utils.CreateResponse(err, "Internal error", false))
	}

	validate := validator.New()
	err := validate.Struct(payload)

	if err != nil {
		fmt.Println("ERR handlers/administration/userDocumentType.go: Updating - Payload validation ", err)
		validationErrors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(nil, validationErrors.Error(), false))
	}

	identifier := c.Param("identifier")
	currentModel := &models.UserDocumentType{}
	currentModel.ID = identifier

	currentModel, err = currentModel.QueryByID(db)
	if err != nil {
		fmt.Println("ERR handlers/administration/userDocumentType.go: Updating - getting model ", err)
		return c.JSON(http.StatusNotFound, utils.CreateResponse(nil, "User Document Type not found", true))
	}

	// New fields

	currentModel.Label = payload.Label
	currentModel.Abbreviation = payload.Abbreviation

	dataDB, err := currentModel.Update(db)
	if err != nil {
		fmt.Println("ERR handlers/administration/userDocumentType.go: Updating - updating model ", err)
		return utils.ReturnError(err, c)
	}

	return c.JSON(http.StatusOK, utils.CreateResponse(dataDB, "User Document type updated", true))

}
