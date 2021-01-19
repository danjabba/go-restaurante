package models

// Food model ...
type Food struct {
	Base
	Name        string   `gorm:"type:varchar(20)" json:"name"`
	Description string   `gorm:"type:varchar(50)" json:"description"`
	Ingredients []string `gorm:"type:varchar(50)" json:"ingredients"`
	Price       float64  `gorm:"type:varchar(15)" json:"price"`
}
