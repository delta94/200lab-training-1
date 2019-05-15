package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

// Account is struct contain a infor related to a user
type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

// Token is struct contain a token related to a user
type Token struct {
	UserID uint
	jwt.StandardClaims
}
