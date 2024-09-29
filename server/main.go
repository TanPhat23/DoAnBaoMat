package main

import (
	"context"
	auth "doAnBaoMat/middleware"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
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

type User struct {
	IDDB     primitive.ObjectID `bson:"_id"`
	Id       string
	Username string
	Password string
	Role     string
}


/*global variable*/

var secretKey = []byte(os.Getenv("SECRET_KEY"))
var loggedInUser User
var todos []Todo

/* MONGO */
func getMongoUser(username string, password string) bool {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		return false
	}
	coll := client.Database("DoAnBaoMat").Collection("Users")
	name := username

	err = coll.FindOne(context.TODO(), bson.D{{"Name", name}}).Decode(&loggedInUser)
	if err != nil {
		panic(err)
	}

	if password == result["Password"] {
		loggedInUser.user=result
		return true
	}
	return false
}

func getTodos() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		return
	}
	coll := client.Database("DoAnBaoMat").Collection("Todos")

	cursor, err := coll.Find(context.TODO(), bson.D{})

	if err = cursor.All(context.TODO(), &todos); err != nil {
		panic(err)
	}
}

func addToDoToDB(todo Todo) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		return
	}
	coll := client.Database("DoAnBaoMat").Collection("Todos")

	result, err := coll.InsertOne(context.TODO(), todo)

	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted document %v", result)
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
		c.SetCookie("token", tokenString, 3600, "/", "", true, true)
		c.IndentedJSON(http.StatusOK, tokenString)
	} else {
		c.String(http.StatusUnauthorized, "Invalid credentials")
	}
}

func getData(c *gin.Context) {
	getTodos()
	c.IndentedJSON(http.StatusOK, todos)
}

func getCurrentUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, loggedInUser)
	if loggedInUser.Username == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}
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
	c.SetCookie("token", "", -1, "/", "", true, true)
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
