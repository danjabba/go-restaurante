package models

import "gorm.io/gorm"

// Restaurant model ...
type Restaurant struct {
	Base
	Name    string `gorm:"type:varchar(25)" json:"name"`
	AdminID string `gorm:"unique;type:varchar(50)" json:"admin_id"`
	Admin   User   `gorm:"foreignKey:AdminID" json:"admin"`
	Address string `gorm:"type:varcha(25)" json:"address"`
}

// TableName ...
func (*Restaurant) TableName() string {

	return "restaurants"
}

// Create ...
func (c *Restaurant) Create(conn *gorm.DB) error {

	result := conn.Create(c)
	if result.Error != nil {
		return result.Error
	}

	return nil

}

// Delete ...
func (c *Restaurant) Delete(conn *gorm.DB) (*Restaurant, error) {

	restaurant := &Restaurant{}

	result := conn.Delete(restaurant, c)

	if result.Error != nil {
		return restaurant, result.Error
	}

	return restaurant, nil

}

// Query structure method.
func (c *Restaurant) Query(conn *gorm.DB) ([]*Restaurant, error) {

	var restaurant []*Restaurant

	result := conn.Where(c).Find(&restaurant)
	if result.Error != nil {
		return restaurant, result.Error
	}

	return restaurant, nil

}

// QueryByID ...
func (c *Restaurant) QueryByID(conn *gorm.DB) (*Restaurant, error) {

	restaurant := &Restaurant{}

	result := conn.Where("id = ?", c.ID).First(&restaurant)

	if result.Error != nil {
		return restaurant, result.Error
	}

	return restaurant, nil

}

// Update ...
func (c *Restaurant) Update(conn *gorm.DB) (*Restaurant, error) {

	result := conn.Save(c)
	if result.Error != nil {
		return c, result.Error
	}

	return c, nil
}
