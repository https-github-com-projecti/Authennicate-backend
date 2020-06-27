package repository

import (
	"authenName/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUserById(id string) (model.User, error)
	CreateUser(user model.User) error
	Login(username string, password string) (model.User, error)
}

var userCollection *mongo.Collection

func CreateCollectionUser(DB *mongo.Database) {
	userCollection = DB.Collection("user")
}

func Login(username string, password string) (model.User, error) {
	user := model.User{}
	err := userCollection.FindOne(context.TODO(), bson.M{"username": username, "password": password}).Decode(&user)
	return user, err
}

func GetUserById(id string) (model.User, error) {
	user := model.User{}
	docID, err := primitive.ObjectIDFromHex(id)
	err = userCollection.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&user)
	return user, err
}

func CreateUser(user model.User) error {
	_, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
		return err
	}
	return nil
}
