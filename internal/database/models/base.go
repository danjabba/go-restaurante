package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// Base contains common columns for all tables.
type Base struct {
	ID        string     `gorm:"primary_key;" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (u *Base) BeforeCreate(tx *gorm.DB) error {

	u.ID = uuid.NewV4().String()

	return nil

}
