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

// createRestaurantPayload ...
type createRestaurantPayload struct {
	Name    string `bson:"name" json:"name"`
	AdminID string `bson:"admin_id" validate:"required" json:"admin_id"`
	Address string `bson:"address" json:"address"`
}

// updateRestaurantPayload ...
type updateRestaurantPayload struct {
	Name    string `bson:"name" json:"name"`
	AdminID string `bson:"admin_id" validate:"required" json:"admin_id"`
	Address string `bson:"address" json:"address"`
}

// ListRestaurant ...
func ListRestaurant(c echo.Context) error {

	db := tools.GetConnection()

	model := &models.Restaurant{}
	dataDB, err := model.Query(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/restaurant.go: Error listing", err)
		return utils.ReturnError(err, c)
	}

	return c.JSON(http.StatusOK, utils.CreateResponse(dataDB, "restaurants are listed", true))

}

// CreateRestaurant ...
func CreateRestaurant(c echo.Context) error {

	db := tools.GetConnection()

	payload := new(createRestaurantPayload)

	if err := c.Bind(payload); err != nil {

		fmt.Println("ERR handlers/administration/restaurant.go: Creating - Binding payload", err)
		return c.JSON(http.StatusUnprocessableEntity, utils.CreateResponse(err, "internal error", false))

	}

	validate := validator.New()
	err := validate.Struct(payload)

	if err != nil {

		fmt.Println("ERR handlers/administration/restaurant.go: Creating - Payload validation", err)
		validationerrors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(nil, validationerrors.Error(), false))

	}

	model := &models.Restaurant{
		Name:    payload.Name,
		AdminID: payload.AdminID,
		Address: payload.Address,
	}

	err = model.Create(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/restaurant.go: Creating - Saving model", err)
		return utils.ReturnError(err, c)
	}

	return c.JSON(http.StatusCreated, utils.CreateResponse(model, "restaurant created", true))

}

// DeleteRestaurant ...
func DeleteRestaurant(c echo.Context) error {

	db := tools.GetConnection()

	identifier := c.Param("identifier")

	model := &models.Restaurant{}
	model.ID = identifier

	_, err := model.QueryByID(db)

	if err != nil {

		return c.JSON(http.StatusNotFound, utils.CreateResponse(nil, "restaurant not found", true))
	}

	dataDB, err := model.Delete(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/restaurant.go: Deleting - Deleting model", err)
		return utils.ReturnError(err, c)
	}

	return c.JSON(http.StatusOK, utils.CreateResponse(dataDB, "restaurant deleted", true))

}

// GetRestaurant ...
func GetRestaurant(c echo.Context) error {

	db := tools.GetConnection()

	identifier := c.Param("identifier")

	model := &models.Restaurant{}
	model.ID = identifier

	dataDB, err := model.QueryByID(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/restaurant.go: Getting - getting model: ", err)
		return c.JSON(http.StatusNotFound, utils.CreateResponse(nil, "restaurant not found", true))
	}

	return c.JSON(http.StatusOK, utils.CreateResponse(dataDB, "restaurant retrieved", true))

}

// UpdateRestaurant ...
func UpdateRestaurant(c echo.Context) error {

	db := tools.GetConnection()

	payload := new(updateRestaurantPayload)

	if err := c.Bind(payload); err != nil {

		fmt.Println("ERR handlers/administration/restaurant.go: Updating - Binding payload ", err)
		return c.JSON(http.StatusUnprocessableEntity, utils.CreateResponse(err, "internal error", false))

	}

	validate := validator.New()
	err := validate.Struct(payload)

	if err != nil {

		fmt.Println("ERR handlers/administration/restaurant.go: Updating - Payload validation ", err)
		validationErrors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(nil, validationErrors.Error(), false))
	}

	identifier := c.Param("identifier")
	currentModel := &models.Restaurant{}
	currentModel.ID = identifier

	currentModel, err = currentModel.QueryByID(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/restaurant.go: Updating - getting model", err)
		return c.JSON(http.StatusNotFound, utils.CreateResponse(nil, "restaurant not found", true))

	}

	// New fields

	currentModel.Name = payload.Name
	currentModel.AdminID = payload.AdminID
	currentModel.Address = payload.Address

	dataDB, err := currentModel.Update(db)
	if err != nil {

		fmt.Println("ERR handlers/administration/restaurant.go: Updating - updating model", err)
		return utils.ReturnError(err, c)

	}

	return c.JSON(http.StatusOK, utils.CreateResponse(dataDB, "restaurant Updated", true))
}
