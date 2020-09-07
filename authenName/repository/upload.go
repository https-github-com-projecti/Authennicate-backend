package repository

import (
	"authenName/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type UploadRepository interface {
	InsertUpload(upload model.Upload) (*mongo.InsertOneResult, error)
	FindByPathAndStatus(id primitive.ObjectID) (model.Upload, error)
	DeleteById(id string) (*mongo.DeleteResult, error)
	FindById(id string) (model.Upload, error)
	FindAllByStatus(status string) ([]*model.Upload, error)
	FindByUserId(id string) (model.Upload, error)
}

var uploadCollection *mongo.Collection

func CreateCollectionUpload(DB *mongo.Database) {
	uploadCollection = DB.Collection("upload")
}

func InsertUpload(upload model.Upload) (*mongo.InsertOneResult, error) {
	insertResult , err := uploadCollection.InsertOne(context.TODO(), upload)
	if err != nil {
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

func FindAllByStatus(status string) ([]*model.Upload, error){
	var upload []*model.Upload
	res, errFind := uploadCollection.Find(context.TODO(), bson.M{"status": status})
	if errFind != nil {
		log.Fatal(errFind)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for res.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem model.Upload
		err := res.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		upload = append(upload, &elem)
	}

	if err := res.Err(); err != nil {
		log.Fatal(err)
	}
	res.Close(context.TODO())
	return upload, errFind
}

func FindByUserId(id string) (model.Upload, error){
	upload := model.Upload{}
	docID, err := primitive.ObjectIDFromHex(id)
	err = uploadCollection.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&upload)
	return upload, err
}
