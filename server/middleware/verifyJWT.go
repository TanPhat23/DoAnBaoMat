package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)
var secretKey []byte

func VerifyToken(tokenString string) (*jwt.Token, error){
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil{
		return nil, err
	}
	if !token.Valid{
		return nil, fmt.Errorf("Invalid token")
	}
	return token, nil
}

func AuthenticateMiddleware(c *gin.Context){
	tokenString, err := c.Cookie("access_token")
	if err != nil{
		c.Abort()
		return
	}
	
	token, err := VerifyToken(tokenString)
	if err != nil{
		c.Abort()
		return
	}
	fmt.Printf("Token verified successfully. Claims: %+v\\n", token.Claims)
	c.Next()
}