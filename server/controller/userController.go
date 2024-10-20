package controller

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"

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


func RefreshMongoUser(user * User) User{
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		panic(err)
	}
	coll := client.Database("DoAnBaoMat").Collection("Users")
	var dataBaseUser User
	err = coll.FindOne(context.TODO(), bson.D{{"Name", user.Username}}).Decode(&dataBaseUser)
	if err != nil {
		panic(err)
	}
	return dataBaseUser
}

func GetMongoUser(user *User) User {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		panic(err)
	}
	coll := client.Database("DoAnBaoMat").Collection("Users")
	name := user.Username
	
	var dataBaseUser User
	err = coll.FindOne(context.TODO(), bson.D{{"Name", name}}).Decode(&dataBaseUser)
	if err != nil {
		panic(err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(dataBaseUser.Password), []byte(user.Password)); err != nil {
		panic(err)
	}
	return dataBaseUser
}

func CreateMongoUser(user *User) (User, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		panic(err)
	}
	//Check if user exist or not
	var existUser User
	coll := client.Database("DoAnBaoMat").Collection("Users")
	err = coll.FindOne(context.TODO(), bson.D{{"Name", user.Username}}).Decode(&existUser)

	if existUser.Username == user.Username {
		panic(err)
	} else {
		encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
		if err != nil {
			panic(err)
		}

		newUser := &User{
			IDDB:     primitive.NewObjectID(),
			Id:       strconv.Itoa(rand.Int()),
			Username: user.Username,
			Password: string(encryptedPassword),
			Role:     "employee",
		}

		result, err := coll.InsertOne(context.TODO(), newUser)

		if err != nil {
			panic(err)
		}
		fmt.Println("Inserted document %v", result)
		return *newUser, nil
	}
}
