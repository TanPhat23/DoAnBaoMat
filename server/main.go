package main

import (
	controller "doAnBaoMat/controller"
	auth "doAnBaoMat/middleware"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

/*global variable*/

/* JWT */

/* HTTP */
func login(c *gin.Context) {
	var user controller.User
	if err := c.BindJSON(&user); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if controller.GetMongoUser(&user) == true {
		controller.IssueTokens(c, user)
		fmt.Printf("Token created")
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid credential"})
	}
}

func signin(c *gin.Context) {
	var user controller.User
	if err := c.BindJSON(&user); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := controller.CreateMongoUser(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User already exist"})
	} else {
		controller.IssueTokens(c, user)
		fmt.Printf("Token created")
	}
}

func getData(c *gin.Context) {
	controller.GetTodos()
	c.IndentedJSON(http.StatusOK, controller.Todos)
}

func getCurrentUser(c *gin.Context) {
	tokenString, err := c.Cookie("access_token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	token, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}
	user := controller.User{
		Username: claims["Username"].(string),
		Role:     claims["Role"].(string),
	}
	c.IndentedJSON(http.StatusOK, user)
}

func addToDo(c *gin.Context) {
	var todo controller.Todo
	if err := c.BindJSON(&todo); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	controller.AddToDoToDB(todo)
	c.String(http.StatusOK, "Todo added")
}

func logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, false)
}

/* main function */
func main() {

	godotenv.Load()
	router := gin.Default()

	//CROSS ORIGIN ACCESS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           3600,
	}))

	router.Static("/static", "./static")

	/* END POINTS*/
	router.GET("/todos", auth.AuthenticateMiddleware, getData)
	router.GET("/currentuser", auth.AuthenticateMiddleware, getCurrentUser)

	router.POST("/add", auth.AuthenticateMiddleware, addToDo)
	router.POST("/login", login)
	router.POST("/signin", signin)

	router.DELETE("/logout", logout)

	router.Run(":8080")
}
