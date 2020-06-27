package repository

import (
	"authenName/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UploadRepository interface {
	InsertUpload(upload model.Upload) (*mongo.InsertOneResult, error)
	FindByPathAndStatus(id primitive.ObjectID) (model.Upload, error)
	DeleteById(id string) (*mongo.DeleteResult, error)
}

var uploadCollection *mongo.Collection

func CreateCollectionUpload(DB *mongo.Database) {
	uploadCollection = DB.Collection("upload")
}

func InsertUpload(upload model.Upload) (*mongo.InsertOneResult, error) {
	insertResult , err := uploadCollection.InsertOne(context.TODO(), upload)
	if err != nil {
		panic(err)
		return nil, err
	}
	return insertResult, nil
}

func FindByPathAndStatus(path string, status string) (model.Upload, error) {
	upload := model.Upload{}
	err := uploadCollection.FindOne(context.TODO(), bson.M{"path": path, "status": status}).Decode(&upload)
	return upload, err
}

func DeleteById(id string) (*mongo.DeleteResult, error) {
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	result, errDel := uploadCollection.DeleteOne(context.TODO(), bson.M{"_id": docID})
	if errDel != nil {
		panic(errDel)
	}
	return result, errDel
}

func FindById(id string) (model.Upload, error) {
	uploadFile := model.Upload{}
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	errFile := uploadCollection.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&uploadFile)
	return uploadFile, errFile
}
