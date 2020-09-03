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
	DeleteUserById(id string) (*mongo.DeleteResult, error)
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

func DeleteUserById(id string) (*mongo.DeleteResult, error) {
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	result, errDel := userCollection.DeleteOne(context.TODO(), bson.M{"_id": docID})
	if errDel != nil {
		panic(errDel)
	}
	return result, errDel
}

func UpdateUserById(user model.User) error {
	_, err := userCollection.UpdateOne(context.TODO(), bson.M{"_id": user.ID}, bson.M{"$set": bson.M{
		"createUpdate": user.CreateUpdate,
		"username":     user.Username,
		"phoneNumber":  user.PhoneNumber,
		"email":        user.Email,
	}})
	if err != nil {
		panic(err)
		return err
	}
	return nil
}
