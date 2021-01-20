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

// createDrinkPayload ...
type createDrinkPayload struct {
	Name        string   `bson:"name" json:"name"`
	Description string   `bson:"description" json:"description"`
	Ingredients []string `bson:"ingredients" json:"ingredients"`
	Price       float64  `bson:"price" json:"price"`
}

// updateDrinkPayload ...
type updateDrinkPayload struct {
	Name        string   `bson:"name" json:"name"`
	Description string   `bson:"description" json:"description"`
	Ingredients []string `bson:"ingredients" json:"ingredients"`
	Price       float64  `bson:"price" json:"price"`
}

// ListDrink ...
func ListDrink(c echo.Context) error {

	db := tools.GetConnection()

	model := &models.Drink{}
	dataDB, err := model.Query(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/drink.go: Error listing", err)
		return utils.ReturnError(err, c)
	}

	return c.JSON(http.StatusOK, utils.CreateResponse(dataDB, "drinks are listed", true))

}

// CreateDrink ...
func CreateDrink(c echo.Context) error {

	db := tools.GetConnection()

	payload := new(createDrinkPayload)

	if err := c.Bind(payload); err != nil {

		fmt.Println("ERR handlers/administration/drink.go: Creating - Binding payload", err)
		return c.JSON(http.StatusUnprocessableEntity, utils.CreateResponse(err, "internal error", false))

	}

	validate := validator.New()
	err := validate.Struct(payload)

	if err != nil {

		fmt.Println("ERR handlers/administration/drink.go: Creating - Payload validation", err)
		validationerrors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(nil, validationerrors.Error(), false))

	}

	model := &models.Drink{
		Name:        payload.Name,
		Description: payload.Description,
		Ingredients: payload.Ingredients,
		Price:       payload.Price,
	}

	err = model.Create(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/drink.go: Creating - Saving model", err)
		return utils.ReturnError(err, c)
	}

	return c.JSON(http.StatusCreated, utils.CreateResponse(model, "drink created", true))

}

// DeleteDrink ...
func DeleteDrink(c echo.Context) error {

	db := tools.GetConnection()

	identifier := c.Param("identifier")

	model := &models.Drink{}
	model.ID = identifier

	_, err := model.QueryByID(db)

	if err != nil {

		return c.JSON(http.StatusNotFound, utils.CreateResponse(nil, "drink not found", true))
	}

	dataDB, err := model.Delete(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/drink.go: Deleting - Deleting model", err)
		return utils.ReturnError(err, c)
	}

	return c.JSON(http.StatusOK, utils.CreateResponse(dataDB, "drink deleted", true))

}

// GetDrink ...
func GetDrink(c echo.Context) error {

	db := tools.GetConnection()

	identifier := c.Param("identifier")

	model := &models.Drink{}
	model.ID = identifier

	dataDB, err := model.QueryByID(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/drink.go: Getting - getting model: ", err)
		return c.JSON(http.StatusNotFound, utils.CreateResponse(nil, "drink not found", true))
	}

	return c.JSON(http.StatusOK, utils.CreateResponse(dataDB, "drink retrieved", true))

}

// UpdateDrink ...
func UpdateDrink(c echo.Context) error {

	db := tools.GetConnection()

	payload := new(updateDrinkPayload)

	if err := c.Bind(payload); err != nil {

		fmt.Println("ERR handlers/administration/drink.go: Updating - Binding payload ", err)
		return c.JSON(http.StatusUnprocessableEntity, utils.CreateResponse(err, "internal error", false))

	}

	validate := validator.New()
	err := validate.Struct(payload)

	if err != nil {

		fmt.Println("ERR handlers/administration/drink.go: Updating - Payload validation ", err)
		validationErrors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(nil, validationErrors.Error(), false))
	}

	identifier := c.Param("identifier")
	currentModel := &models.Drink{}
	currentModel.ID = identifier

	currentModel, err = currentModel.QueryByID(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/drink.go: Updating - getting model", err)
		return c.JSON(http.StatusNotFound, utils.CreateResponse(nil, "drink not found", true))

	}

	// New fields

	currentModel.Name = payload.Name
	currentModel.Description = payload.Description
	currentModel.Ingredients = payload.Ingredients
	currentModel.Price = payload.Price

	dataDB, err := currentModel.Update(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/drink.go: Updating - updating model", err)
		return utils.ReturnError(err, c)

	}

	return c.JSON(http.StatusOK, utils.CreateResponse(dataDB, "drink Updated", true))
}
