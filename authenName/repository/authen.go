package repository

import (
	"authenName/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type AuthenRepository interface {
	CreateAuthen(authen model.Authen) error
	FindAuthenBySubjectId(id string) ([]*model.Authen, error)
}

var AuthenCollection *mongo.Collection

func CreateCollectionAuthen(DB *mongo.Database) {
	AuthenCollection = DB.Collection("authen")
}

func CreateAuthen(authen model.Authen) error {
	_, err := AuthenCollection.InsertOne(context.TODO(), authen)
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func FindAuthenBySubjectId(id string) ([]*model.Authen, error) {
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	var authen []*model.Authen
	res, errFind := AuthenCollection.Find(context.TODO(), bson.M{"Subject": docID})
	if errFind != nil {
		log.Fatal(errFind)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for res.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem model.Authen
		err := res.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		authen = append(authen, &elem)
	}

	if err := res.Err(); err != nil {
		log.Fatal(err)
	}
	res.Close(context.TODO())
	return authen, errFind
}
