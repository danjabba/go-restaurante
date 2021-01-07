package models

import (
	"gorm.io/gorm"
)

// UserDocumentType model ...
type UserDocumentType struct {
	Base
	Label        string `gorm:"unique:idx_label,type:varchar(25)" json:"label"`
	Abbreviation string `gorm:"unique:idx_abbreviation,type:varchar(25)" json:"abbreviation"`
}

// TableName ...
func (UserDocumentType) TableName() string {
	return "user_document_types"
}

// Create structure method.
func (base *UserDocumentType) Create(conn *gorm.DB) error {

	result := conn.Create(base)
	if result.Error != nil {
		return result.Error
	}

	return nil

}

// Query structure method.
func (base *UserDocumentType) Query(conn *gorm.DB) ([]*UserDocumentType, error) {

	var userDocumentType []*UserDocumentType

	result := conn.Where(base).Find(&userDocumentType)
	if result.Error != nil {
		return userDocumentType, result.Error
	}

	return userDocumentType, nil

}

// QueryByID by ID.
func (base *UserDocumentType) QueryByID(conn *gorm.DB) (*UserDocumentType, error) {

	userDocumentType := &UserDocumentType{}

	result := conn.Where("id = ?", base.ID).First(&userDocumentType)

	if result.Error != nil {
		return userDocumentType, result.Error
	}

	return userDocumentType, nil

}

// Update ...
func (base *UserDocumentType) Update(conn *gorm.DB) (*UserDocumentType, error) {

	result := conn.Save(base)
	if result.Error != nil {
		return base, result.Error
	}

	return base, nil

}

// Delete ...
func (base *UserDocumentType) Delete(conn *gorm.DB) (*UserDocumentType, error) {

	userDocumentType := &UserDocumentType{}

	result := conn.Delete(userDocumentType, base)

	if result.Error != nil {
		return userDocumentType, result.Error
	}

	return userDocumentType, nil

}
