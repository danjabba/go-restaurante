package models

import "gorm.io/gorm"

// Drink model ...
type Drink struct {
	Base
	Name        string   `gorm:"type:varchar(20)" json:"name"`
	Description string   `gorm:"type:varchar(50)" json:"description"`
	Ingredients []string `gorm:"type:varchar(50)" json:"ingredients"`
	Price       float64  `gorm:"type:varchar(15)" json:"price"`
}

// TableName ...
func (*Drink) TableName() string {

	return "drinks"
}

// Create ...
func (c *Drink) Create(conn *gorm.DB) error {

	result := conn.Create(c)
	if result.Error != nil {
		return result.Error
	}

	return nil

}

// Delete ...
func (c *Drink) Delete(conn *gorm.DB) (*Drink, error) {

	drink := &Drink{}

	result := conn.Delete(drink, c)

	if result.Error != nil {
		return drink, result.Error
	}

	return drink, nil

}

// Query structure method.
func (c *Drink) Query(conn *gorm.DB) ([]*Drink, error) {

	var drink []*Drink

	result := conn.Where(c).Find(&drink)
	if result.Error != nil {
		return drink, result.Error
	}

	return drink, nil

}

// QueryByID ...
func (c *Drink) QueryByID(conn *gorm.DB) (*Drink, error) {

	drink := &Drink{}

	result := conn.Where("id = ?", c.ID).First(&drink)

	if result.Error != nil {
		return drink, result.Error
	}

	return drink, nil

}

// Update ...
func (c *Drink) Update(conn *gorm.DB) (*Drink, error) {

	result := conn.Save(c)
	if result.Error != nil {
		return c, result.Error
	}

	return c, nil
}
