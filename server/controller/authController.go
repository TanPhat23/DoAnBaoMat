package controller

import (
	auth "doAnBaoMat/middleware"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

func IssueTokens(c *gin.Context, user User) {
	accessToken, err := generateAccessToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	refreshToken, err := generateRefreshToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}

	c.SetCookie("access_token", accessToken, 3600, "/", "", false, true)
	c.SetCookie("refresh_token", refreshToken, 3600*24*7, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"access_token": accessToken, "refresh_token": refreshToken})
}

func generateAccessToken(user User) (string, error) {
	claims := jwt.MapClaims{
		"username": user.Username,
		"iss":      "todo-app",
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func generateRefreshToken(user User) (string, error) {
	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func RefreshToken(c *gin.Context) {
    refreshToken, err := c.Cookie("refresh_token")
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not provided"})
        return
    }

    token, err := auth.VerifyToken(refreshToken)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
        return
    }

    claims := token.Claims.(jwt.MapClaims)
    username := claims["username"].(string)
    user := RefreshMongoUser(&User{Username: username})

    IssueTokens(c, user)
}