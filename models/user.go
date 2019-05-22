package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User is struct contain a infor related to a user
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null",binding:"required"`
	Token    string `gorm:"unique;"`
	Password string `binding:"required"`
}

type Login struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

type SignUpResponse struct {
	ID       uint
	Username string
	Email    string
	Fullname string
	Bod      *time.Time
}

type LogInResponse struct {
	ID       uint
	Fullname string
	Token    string
}
