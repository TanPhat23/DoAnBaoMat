package main

import (
	auth "doAnBaoMat/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Todo struct {
	Text string
	Done bool
}

type User struct {
	Id       string
	UserName string
	Password string
	Role     string
}

/*global variable*/
var secretKey = []byte(os.Getenv("SECRET_KEY"))
var loggedInUser string
var todos []Todo
var users []User = []User{{
	Id: "1", UserName: "Phat", Password: "password", Role: "senior",
}, {
	Id: "2", UserName: "Luan", Password: "password", Role: "employee",
}, {
	Id: "3", UserName: "Thong", Password: "password", Role: "employee",
}, {
	Id: "4", UserName: "Khoa", Password: "password", Role: "employee",
},
}

func toggleIndex(index string) {
	i, _ := strconv.Atoi(index)
	if i >= 0 && i < len(todos) {
		todos[i].Done = !todos[i].Done
	}

}

func getRole(username string) string {
	if username == "senior" {
		return "senior"
	}
	return "employee"
}

/* JWT */

func createToken(username string) (string, error) {
	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iss": "todo-app",
		"aud": getRole(username),
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
	username := c.PostForm("username")
	password := c.PostForm("password")

	if (username == "employee" && password == "password") || (username == "senior" && password == "password") {
		tokenString, err := createToken(username)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error creating token")
			return
		}

		loggedInUser = username
		fmt.Printf("Token created")
		c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
		c.Redirect(http.StatusSeeOther, "/")
	} else {
		c.String(http.StatusUnauthorized, "Invalid credentials")
	}
}

func getData(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Todos":    todos,
		"LoggedIn": loggedInUser != "",
		"UserName": loggedInUser,
		"Role":     getRole(loggedInUser),
	})
}

func addToDo(c *gin.Context) {
	text := c.PostForm("todo")
	todo := Todo{Text: text, Done: false}
	todos = append(todos, todo)
	c.Redirect(http.StatusSeeOther, "/")
}

func toggleForm(c *gin.Context) {
	index := c.PostForm("index")
	toggleIndex(index)
	c.Redirect(http.StatusSeeOther, "/")
}

func logout(c *gin.Context) {
	loggedInUser = ""
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.Redirect(http.StatusUnauthorized, "/")
}

/* main function */
func main() {
	router := gin.Default()
	router.Static("/static", "./static")

	router.LoadHTMLGlob("templates/*")
	/* END POINTS*/
	router.GET("/", getData)

	router.POST("/login", login)
	router.POST("/logout", logout)

	router.POST("/add", auth.AuthenticateMiddleware, addToDo)
	router.POST("/toggle", auth.AuthenticateMiddleware, toggleForm)

	router.Run(":8080")
}
