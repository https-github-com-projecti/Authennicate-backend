package repository

import (
	"authenName/model"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoleRepository interface {
	GetRoleAll() ([]*model.Role, error)
	CreateRole(role model.Role) error
	GetRoleWithIDUser(id string) (model.Role, error)
}

var roleCollection *mongo.Collection

func CreateCollectionRole(DB *mongo.Database) {
	roleCollection = DB.Collection("role")
}

func GetRoleAll() ([]*model.Role, error) {
	var role []*model.Role
	res, err := roleCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}

	for res.Next(context.TODO()) {
		var elem model.Role
		err := res.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		role = append(role, &elem)
	}

	if err := res.Err(); err != nil {
		log.Fatal(err)
	}
	res.Close(context.TODO())
	return role, err
}

func CreateRole(role model.Role) error {
	_, err := roleCollection.InsertOne(context.TODO(), role)
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func GetRoleWithIDUser(id string) (model.Role, error) {
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	role := model.Role{}
	err = roleCollection.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&role)
	return role, err
}
