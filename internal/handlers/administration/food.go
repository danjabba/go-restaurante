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

// createFoodPayload ...
type createFoodPayload struct {
	Name        string   `bson:"name" json:"name"`
	Description string   `bson:"description" json:"description"`
	Ingredients []string `bson:"ingredients" json:"ingredients"`
	Price       float64  `bson:"price" json:"price"`
}

// updateFoodPayload ...
type updateFoodPayload struct {
	Name        string   `bson:"name" json:"name"`
	Description string   `bson:"description" json:"description"`
	Ingredients []string `bson:"ingredients" json:"ingredients"`
	Price       float64  `bson:"price" json:"price"`
}

// ListFood ...
func ListFood(c echo.Context) error {

	db := tools.GetConnection()

	model := &models.Food{}
	dataDB, err := model.Query(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/food.go: Error listing", err)
		return utils.ReturnError(err, c)
	}

	return c.JSON(http.StatusOK, utils.CreateResponse(dataDB, "foods are listed", true))

}

// CreateFood ...
func CreateFood(c echo.Context) error {

	db := tools.GetConnection()

	payload := new(createFoodPayload)

	if err := c.Bind(payload); err != nil {

		fmt.Println("ERR handlers/administration/food.go: Creating - Binding payload", err)
		return c.JSON(http.StatusUnprocessableEntity, utils.CreateResponse(err, "internal error", false))

	}

	validate := validator.New()
	err := validate.Struct(payload)

	if err != nil {

		fmt.Println("ERR handlers/administration/food.go: Creating - Payload validation", err)
		validationerrors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(nil, validationerrors.Error(), false))

	}

	model := &models.Food{
		Name:        payload.Name,
		Description: payload.Description,
		Ingredients: payload.Ingredients,
		Price:       payload.Price,
	}

	err = model.Create(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/food.go: Creating - Saving model", err)
		return utils.ReturnError(err, c)
	}

	return c.JSON(http.StatusCreated, utils.CreateResponse(model, "food created", true))

}

// DeleteFood ...
func DeleteFood(c echo.Context) error {

	db := tools.GetConnection()

	identifier := c.Param("identifier")

	model := &models.Food{}
	model.ID = identifier

	_, err := model.QueryByID(db)

	if err != nil {

		return c.JSON(http.StatusNotFound, utils.CreateResponse(nil, "food not found", true))
	}

	dataDB, err := model.Delete(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/food.go: Deleting - Deleting model", err)
		return utils.ReturnError(err, c)
	}

	return c.JSON(http.StatusOK, utils.CreateResponse(dataDB, "food deleted", true))

}

// GetFood ...
func GetFood(c echo.Context) error {

	db := tools.GetConnection()

	identifier := c.Param("identifier")

	model := &models.Food{}
	model.ID = identifier

	dataDB, err := model.QueryByID(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/food.go: Getting - getting model: ", err)
		return c.JSON(http.StatusNotFound, utils.CreateResponse(nil, "food not found", true))
	}

	return c.JSON(http.StatusOK, utils.CreateResponse(dataDB, "food retrieved", true))

}

// UpdateFood ...
func UpdateFood(c echo.Context) error {

	db := tools.GetConnection()

	payload := new(updateFoodPayload)

	if err := c.Bind(payload); err != nil {

		fmt.Println("ERR handlers/administration/food.go: Updating - Binding payload ", err)
		return c.JSON(http.StatusUnprocessableEntity, utils.CreateResponse(err, "internal error", false))

	}

	validate := validator.New()
	err := validate.Struct(payload)

	if err != nil {

		fmt.Println("ERR handlers/administration/food.go: Updating - Payload validation ", err)
		validationErrors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(nil, validationErrors.Error(), false))
	}

	identifier := c.Param("identifier")
	currentModel := &models.Food{}
	currentModel.ID = identifier

	currentModel, err = currentModel.QueryByID(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/food.go: Updating - getting model", err)
		return c.JSON(http.StatusNotFound, utils.CreateResponse(nil, "food not found", true))

	}

	// New fields

	currentModel.Name = payload.Name
	currentModel.Description = payload.Description
	currentModel.Ingredients = payload.Ingredients
	currentModel.Price = payload.Price

	dataDB, err := currentModel.Update(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/food.go: Updating - updating model", err)
		return utils.ReturnError(err, c)

	}

	return c.JSON(http.StatusOK, utils.CreateResponse(dataDB, "food Updated", true))
}
