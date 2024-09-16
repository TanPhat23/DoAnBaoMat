package auth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)
var secretKey = []byte(os.Getenv("SECRET_KEY"))

func verifyToken(tokenString string) (*jwt.Token, error){
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil{
		return nil, err
	}
	if !token.Valid{
		return nil, fmt.Errorf("invalid toke")
	}
	return token, nil
}

func AuthenticateMiddleware(c *gin.Context){
	tokenString, err := c.Cookie("token")
	if err != nil{
		fmt.Println("Token missing in cookie")
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}
	
	token, err := verifyToken(tokenString)
	if err != nil{
		fmt.Println("Verify token failed")
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}
	fmt.Printf("Token verified successfully. Claims: %+v\\n", token.Claims)
	c.Next()
}