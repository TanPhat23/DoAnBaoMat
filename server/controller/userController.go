package controller

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	IDDB     primitive.ObjectID `bson:"_id"`
	Id       string
	Username string `bson:"Name"`
	Password string `bson:"Password"`
	Role     string
}

var LoggedInUser User

func GetMongoUser(username string, password string) bool {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		return false
	}
	coll := client.Database("DoAnBaoMat").Collection("Users")
	name := username

	err = coll.FindOne(context.TODO(), bson.D{{"Name", name}}).Decode(&LoggedInUser)
	if err != nil {
		panic(err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(LoggedInUser.Password), []byte(password)); err != nil {
		panic(err)
	}
	return true
}

func CreateMongoUser(username string, password string) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		return err
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	newUser := &User{
		Id:       "0",
		Username: username,
		Password: string(encryptedPassword),
		Role:     "employee",
	}

	if err != nil {
		panic(err)
	}

	coll := client.Database("DoAnBaoMat").Collection("Users")

	result, err := coll.InsertOne(context.TODO(), newUser)

	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted document %v", result)
	LoggedInUser = *newUser
	return nil
}
