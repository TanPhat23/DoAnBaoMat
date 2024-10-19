package controller

import (
	"context"
	"fmt"
	"os"

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
var Todos []Todo

func GetTodos() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		return
	}
	coll := client.Database("DoAnBaoMat").Collection("Todos")

	cursor, err := coll.Find(context.TODO(), bson.D{})

	if err = cursor.All(context.TODO(), &Todos); err != nil {
		panic(err)
	}
}

func AddToDoToDB(todo Todo) {
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