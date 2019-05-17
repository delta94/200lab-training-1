package models

import (
	"github.com/jinzhu/gorm"
)

// Note is struct contains all note related to a user
type Note struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	UserID    uint   `json:"user_id"`
}
