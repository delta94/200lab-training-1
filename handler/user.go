package handler

import (
	"strconv"
	"time"

	"github.com/200lab-training-1/models"
	"github.com/200lab-training-1/repositories"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context, userRepo repositories.UserRepo) (*models.SignUpResponse, error) {
	user := models.User{}
	if err := c.ShouldBind(&user); err != nil {
		return nil, err
	}
	password := []byte(user.Password)
	hashPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	user.Password = string(hashPassword)

	createdUser, err := userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	userSignUpResponse := &models.SignUpResponse{
		ID:       createdUser.ID,
		Username: createdUser.Username,
		Email:    createdUser.Email,
		Fullname: createdUser.Fullname,
		Bod:      createdUser.Bod,
	}

	return userSignUpResponse, nil

}

func Login(c *gin.Context, userRepo repositories.UserRepo) (*models.LogInResponse, error) {
	userLogin := models.Login{}
	if err := c.ShouldBind(&userLogin); err != nil {
		return nil, err
	}

	user, err := userRepo.Find(userLogin.Email)
	if err != nil {
		return nil, err
	}

	password := []byte(userLogin.Password)
	err = bcrypt.CompareHashAndPassword([]byte(userLogin.Password), password)
	if err != nil {
		return nil, err
	}

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(),
		Issuer:    "200lab",
		Id:        strconv.Itoa(int(user.ID)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString("200lab")
	userLogInResponse := &models.LogInResponse{
		ID:       user.ID,
		Fullname: user.Fullname,
		Token:    tokenString,
	}
	c.SetCookie("Token", tokenString, 3600*24*365, "/", "", false, true)
	return userLogInResponse, err
}
