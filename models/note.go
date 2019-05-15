package models

import (
	"github.com/jinzhu/gorm"
)

// Note is struct contains all note related to a user
type Note struct {
	gorm.Model
	Title     string `json:"name"`
	Completed bool   `json:"phone"`
	UserID    uint   `json:"user_id"`
}
