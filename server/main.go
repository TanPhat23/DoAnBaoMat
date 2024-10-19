package main

import (
	controller "doAnBaoMat/controller"
	auth "doAnBaoMat/middleware"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

/*global variable*/

var secretKey = []byte(os.Getenv("SECRET_KEY"))

/* JWT */

func createToken(username string) (string, error) {
	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iss": "todo-app",
		"aud": controller.LoggedInUser,
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})
	tokenString, err := claim.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

/* HTTP */
func login(c *gin.Context) {
	var user controller.User
	if err := c.BindJSON(&user); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if controller.GetMongoUser(&user) == true {
		tokenString, err := createToken(user.Username)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error creating token")
			return
		}

		fmt.Printf("Token created")
		c.SetCookie("token", tokenString, 3600, "/", "", true, true)
	} else {
		c.String(http.StatusUnauthorized, "Invalid credentials")
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
		c.String(http.StatusInternalServerError, "User already exist")
	} else {
		tokenString, err := createToken(user.Username)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error creating token")
			return
		}

		fmt.Printf("Token created")
		c.SetCookie("token", tokenString, 3600, "/", "", true, true)
	}
}

func getData(c *gin.Context) {
	controller.GetTodos()
	c.IndentedJSON(http.StatusOK, controller.Todos)
}

func getCurrentUser(c *gin.Context) {
	if controller.LoggedInUser.Username == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.IndentedJSON(http.StatusOK, controller.LoggedInUser)
	}
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
    c.Redirect(http.StatusSeeOther, "/login")

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
