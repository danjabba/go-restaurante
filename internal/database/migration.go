package main

import (
	"fmt"

	models "github.com/danjabba/go-restaurante/internal/database/models"
	"github.com/danjabba/go-restaurante/tools"
)

func main() {

	db := tools.GetConnection()

	err := db.AutoMigrate(&models.Drink{})
	if err != nil {
		fmt.Println(err)
	}
	/*
		err = db.AutoMigrate(&models.Food{})
		if err != nil {
			fmt.Println(err)
		}

		err = db.AutoMigrate(&models.Restaurant{})
		if err != nil {
			fmt.Println(err)
		}

		err = db.AutoMigrate(&models.UserDocumentType{})
		if err != nil {
			fmt.Println(err)
		}

		err = db.AutoMigrate(&models.User{})
		if err != nil {
			fmt.Println(err)
		}

		err = db.AutoMigrate(&models.Menu{})
		if err != nil {
			fmt.Println(err)
		}
	*/
}
