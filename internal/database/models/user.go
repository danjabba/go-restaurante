package models

import "gorm.io/gorm"

// User model ...
type User struct {
	Base
	DniType          string           `json:"dni_type"`
	UserDocumentType UserDocumentType `gorm:"foreingKey:DniType"`
	DniNumber        string           `gorm:"type:varchar(25)" validate:"num"`
	Name             string           `gorm:"type:varchar(25)" json:"name"`
	LastName         string           `gorm:"type:varchar(25)" json:"last_name"`
	Email            string           `gorm:"type:varchar(25)" json:"email"`
}

// TableName ...
func (User) TableName() string {

	return "users"
}

// Create structure method.
func (base *User) Create(conn *gorm.DB) error {

	result := conn.Create(base)
	if result.Error != nil {
		return result.Error
	}

	return nil

}

// Query structure method.
func (base *User) Query(conn *gorm.DB) ([]*User, error) {

	var user []*User

	result := conn.Where(base).Find(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil

}

// QueryByID by ID or Slug ...
func (base *User) QueryByID(conn *gorm.DB) (*User, error) {

	user := &User{}

	result := conn.Where("id = ?", base.ID).First(&user)

	if result.Error != nil {

		return user, result.Error
	}

	return user, nil
}

// Update ...
func (base *User) Update(conn *gorm.DB) (*User, error) {

	result := conn.Save(base)
	if result.Error != nil {
		return base, result.Error
	}

	return base, nil
}

// Delete ...
func (base *User) Delete(conn *gorm.DB) (*User, error) {

	user := &User{}

	result := conn.Delete(user, base)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil

}
