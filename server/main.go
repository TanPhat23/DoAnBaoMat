package main

import (
	"context"
	auth "doAnBaoMat/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID   primitive.ObjectID `bson:"_id"`
	Text string
	Done bool
}

/*global variable*/

var secretKey = []byte(os.Getenv("SECRET_KEY"))
var loggedInUser string
var todos []Todo
var role string = ""

func toggleIndex(index string) {
	i, _ := strconv.Atoi(index)
	if i >= 0 && i < len(todos) {
		todos[i].Done = !todos[i].Done
	}

}

/* MONGO */
func getMongoUser(username string, password string) bool {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		return false
	}
	coll := client.Database("DoAnBaoMat").Collection("Users")
	name := username

	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"Name", name}}).Decode(&result)
	if err != nil {
		panic(err)
	}

	if password == result["Password"] {
		role = result["Role"].(string)
		return true
	}
	return false
}

/* JWT */

func createToken(username string) (string, error) {
	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iss": "todo-app",
		"aud": loggedInUser.Role,
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
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if getMongoUser(user.Username, user.Password) == true {
		tokenString, err := createToken(user.Username)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error creating token")
			return
		}

		fmt.Printf("Token created")
		c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
		c.Redirect(http.StatusSeeOther, "/")
	} else {
		c.String(http.StatusUnauthorized, "Invalid credentials")
	}
}

// func signin(c *gin.Context) {
// 	username := c.PostForm("username")
// 	password := c.PostForm("password")
// 	if len(password) < 10 {
// 		c.String(http.StatusBadRequest, "Length of password need to be longer than 10")
// 	} else {
// 		users = append(users, User{Id: strconv.Itoa(len(users) + 1), UserName: username, Password: password, Role: "employee"})
// 		c.Redirect(http.StatusCreated, "/")
// 	}
// }

func getData(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Todos":    todos,
		"LoggedIn": loggedInUser != "",
		"UserName": loggedInUser,
		"Role":     role,
	})
}

func addToDo(c *gin.Context) {
	var todo Todo
	if err := c.BindJSON(&todo); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	addToDoToDB(todo)
	c.String(http.StatusOK, "Todo added")
}

func logout(c *gin.Context) {
	loggedInUser = ""
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.Redirect(http.StatusUnauthorized, "/")
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

	router.LoadHTMLGlob("templates/*")
	/* END POINTS*/
	router.GET("/todos", auth.AuthenticateMiddleware, getData)
	router.GET("/currentuser", auth.AuthenticateMiddleware, getCurrentUser)

	router.POST("/add", auth.AuthenticateMiddleware, addToDo)
	router.POST("/login", login)
	router.POST("/logout", logout)

	router.Run(":8080")
}
