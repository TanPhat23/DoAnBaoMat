package auth

import (
	"doAnBaoMat/controller"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey []byte

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}
	return token, nil
}

func AuthenticateMiddleware(c *gin.Context) {
	tokenString, err := c.Cookie("access_token")
	if err != nil {
		refreshToken, err := c.Cookie("refresh_token")
		if err != nil || refreshToken == "" {
			panic(err)
		}
		RefreshToken(c)
		tokenString, err = c.Cookie("access_token")
		if err != nil {
			panic(err)
		}
	}

	token, err := VerifyToken(tokenString)
	if err != nil {
		c.Abort()
		return
	}
	fmt.Printf("Token verified successfully. Claims: %+v\\n", token.Claims)
	c.Next()
}

func RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not provided"})
		return
	}

	token, err := VerifyToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	user := controller.RefreshMongoUser(&controller.User{Username: username})

	controller.IssueTokens(c, user)
}
