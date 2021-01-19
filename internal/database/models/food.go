package models

import "gorm.io/gorm"

// Food model ...
type Food struct {
	Base
	Name        string   `gorm:"type:varchar(20)" json:"name"`
	Description string   `gorm:"type:varchar(50)" json:"description"`
	Ingredients []string `gorm:"type:varchar(50)" json:"ingredients"`
	Price       float64  `gorm:"type:varchar(15)" json:"price"`
}

// TableName ...
func (*Food) TableName() string {

	return "foods"
}

// Create ...
func (c *Food) Create(conn *gorm.DB) error {

	result := conn.Create(c)
	if result.Error != nil {
		return result.Error
	}

	return nil

}

// Delete ...
func (c *Food) Delete(conn *gorm.DB) (*Food, error) {

	food := &Food{}

	result := conn.Delete(food, c)

	if result.Error != nil {
		return food, result.Error
	}

	return food, nil

}

// Query structure method.
func (c *Food) Query(conn *gorm.DB) ([]*Food, error) {

	var food []*Food

	result := conn.Where(c).Find(&food)
	if result.Error != nil {
		return food, result.Error
	}

	return food, nil

}

// QueryByID ...
func (c *Food) QueryByID(conn *gorm.DB) (*Food, error) {

	food := &Food{}

	result := conn.Where("id = ?", c.ID).First(&food)

	if result.Error != nil {
		return food, result.Error
	}

	return food, nil

}

// Update ...
func (c *Food) Update(conn *gorm.DB) (*Food, error) {

	result := conn.Save(c)
	if result.Error != nil {
		return c, result.Error
	}

	return c, nil
}
