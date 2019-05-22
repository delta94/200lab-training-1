package handler

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/200lab-training-1/models"
	"github.com/200lab-training-1/repositories"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var key = []byte("200lab")

func SignUp(c *gin.Context, userRepo repositories.UserRepo) (string, error) {
	user := models.User{}
	if err := c.ShouldBind(&user); err != nil {
		return "", err
	}

	password := []byte(user.Password)
	hashPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	user.Password = string(hashPassword)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["client"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()

	tokenString, _ := token.SignedString(key)
	user.Token = tokenString

	createdUser, err := userRepo.Create(user)
	if err != nil {
		return "", err
	}
	return createdUser.Token, nil

}

func Login(c *gin.Context, userRepo repositories.UserRepo) (string, error) {
	userLogin := models.User{}
	if err := c.ShouldBind(&userLogin); err != nil {
		return "", err
	}

	user, err := userRepo.Find(userLogin.Username)
	if err != nil {
		return "", errors.New("Not register yet")
	}

	password := []byte(userLogin.Password)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), password)
	if err != nil {
		return "", err
	}

	return user.Token, nil
}
