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

// createMenuPayload ...
type createMenuPayload struct {
	DrinkID []string `bson:"drink_id" json:"drink_id"`
	FoodID  []string `bson:"food_id" json:"food_id"`
}

// updateMenuPayload ...
type updateMenuPayload struct {
	DrinkID []string `bson:"drink_id" json:"drink_id"`
	FoodID  []string `bson:"food_id" json:"food_id"`
}

// ListMenu ...
func ListMenu(c echo.Context) error {

	db := tools.GetConnection()

	model := &models.Menu{}
	dataDB, err := model.Query(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/menu.go: Error listing", err)
		return utils.ReturnError(err, c)
	}

	return c.JSON(http.StatusOK, utils.CreateResponse(dataDB, "menus are listed", true))

}

// CreateMenu ...
func CreateMenu(c echo.Context) error {

	db := tools.GetConnection()

	payload := new(createMenuPayload)

	if err := c.Bind(payload); err != nil {

		fmt.Println("ERR handlers/administration/menu.go: Creating - Binding payload", err)
		return c.JSON(http.StatusUnprocessableEntity, utils.CreateResponse(err, "internal error", false))

	}

	validate := validator.New()
	err := validate.Struct(payload)

	if err != nil {

		fmt.Println("ERR handlers/administration/menu.go: Creating - Payload validation", err)
		validationerrors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(nil, validationerrors.Error(), false))

	}

	model := &models.Menu{
		FoodID:  payload.FoodID,
		DrinkID: payload.DrinkID,
	}

	err = model.Create(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/menu.go: Creating - Saving model", err)
		return utils.ReturnError(err, c)
	}

	return c.JSON(http.StatusCreated, utils.CreateResponse(model, "menu created", true))

}

// DeleteMenu ...
func DeleteMenu(c echo.Context) error {

	db := tools.GetConnection()

	identifier := c.Param("identifier")

	model := &models.Menu{}
	model.ID = identifier

	_, err := model.QueryByID(db)

	if err != nil {

		return c.JSON(http.StatusNotFound, utils.CreateResponse(nil, "menu not found", true))
	}

	dataDB, err := model.Delete(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/menu.go: Deleting - Deleting model", err)
		return utils.ReturnError(err, c)
	}

	return c.JSON(http.StatusOK, utils.CreateResponse(dataDB, "menu deleted", true))

}

// GetMenu ...
func GetMenu(c echo.Context) error {

	db := tools.GetConnection()

	identifier := c.Param("identifier")

	model := &models.Menu{}
	model.ID = identifier

	dataDB, err := model.QueryByID(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/menu.go: Getting - getting model: ", err)
		return c.JSON(http.StatusNotFound, utils.CreateResponse(nil, "menu not found", true))
	}

	return c.JSON(http.StatusOK, utils.CreateResponse(dataDB, "menu retrieved", true))

}

// UpdateMenu ...
func UpdateMenu(c echo.Context) error {

	db := tools.GetConnection()

	payload := new(updateMenuPayload)

	if err := c.Bind(payload); err != nil {

		fmt.Println("ERR handlers/administration/menu.go: Updating - Binding payload ", err)
		return c.JSON(http.StatusUnprocessableEntity, utils.CreateResponse(err, "internal error", false))

	}

	validate := validator.New()
	err := validate.Struct(payload)

	if err != nil {

		fmt.Println("ERR handlers/administration/menu.go: Updating - Payload validation ", err)
		validationErrors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(nil, validationErrors.Error(), false))
	}

	identifier := c.Param("identifier")
	currentModel := &models.Menu{}
	currentModel.ID = identifier

	currentModel, err = currentModel.QueryByID(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/menu.go: Updating - getting model", err)
		return c.JSON(http.StatusNotFound, utils.CreateResponse(nil, "menu not found", true))

	}

	// New fields

	currentModel.DrinkID = payload.DrinkID
	currentModel.FoodID = payload.FoodID

	dataDB, err := currentModel.Update(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/menu.go: Updating - updating model", err)
		return utils.ReturnError(err, c)

	}

	return c.JSON(http.StatusOK, utils.CreateResponse(dataDB, "menu Updated", true))
}
