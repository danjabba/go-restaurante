package models

import "gorm.io/gorm"

// Menu model ...
type Menu struct {
	Base
	FoodID  []string `gorm:"type:varchar(50)" json:"food_id"`
	Food    []Food   `gorm:"foreignKey:FoodID" json:"food"`
	DrinkID []string `gorm:"type:varchar(50)" json:"drink_id"`
	Drink   []Drink  `gorm:"foreignKey:DrinkID" json:"drink"`
}

// TableName ...
func (*Menu) TableName() string {
	return "menus"
}

// Create ...
func (c *Menu) Create(conn *gorm.DB) error {

	result := conn.Create(c)
	if result.Error != nil {
		return result.Error
	}

	return nil

}

// Delete ...
func (c *Menu) Delete(conn *gorm.DB) (*Menu, error) {

	menu := &Menu{}

	result := conn.Delete(menu, c)

	if result.Error != nil {
		return menu, result.Error
	}

	return menu, nil

}

// Query structure method.
func (c *Menu) Query(conn *gorm.DB) ([]*Menu, error) {

	var menu []*Menu

	result := conn.Where(c).Find(&menu)
	if result.Error != nil {
		return menu, result.Error
	}

	return menu, nil

}

// QueryByID ...
func (c *Menu) QueryByID(conn *gorm.DB) (*Menu, error) {

	menu := &Menu{}

	result := conn.Where("id = ?", c.ID).First(&menu)

	if result.Error != nil {
		return menu, result.Error
	}

	return menu, nil

}

// Update ...
func (c *Menu) Update(conn *gorm.DB) (*Menu, error) {

	result := conn.Save(c)
	if result.Error != nil {
		return c, result.Error
	}

	return c, nil
}
