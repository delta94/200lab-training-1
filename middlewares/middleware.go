package middlewares

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	jwtSecretKey = []byte("ThisIsAVerySecretKey")
	identityKey  = "identity"
)

func AuthenMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authentication")
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecretKey, nil
	})

	if err != nil {
		c.AbortWithStatus(401)
		return
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	fmt.Println("jwt-claims:", claims)
	if ok && claims.Valid() == nil {
		c.Set(identityKey, claims.Id)
		c.Next()
		return
	}
	c.AbortWithStatus(401)
}
