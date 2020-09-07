package repository

import (
	"authenName/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type SubjectRepository interface {
	FindKey(key int32) (model.Subject, error)
	CreateSubject(subject model.Subject) error
	DeleteSubjectById(id primitive.ObjectID) (*mongo.DeleteResult, error)
	GetSubjectForUserId(id primitive.ObjectID) ([]*model.Subject, error)
	GetSubjectById(id string) (model.Subject, error)
}

var subjectCollection *mongo.Collection

func CreateCollectionSubject(DB *mongo.Database) {
	subjectCollection = DB.Collection("subject")
}

func FindKey(key int64) (model.Subject, error) {
	subject := model.Subject{}
	err := subjectCollection.FindOne(context.TODO(), bson.M{"key": key}).Decode(&subject)
	return subject, err
}

func CreateSubject(subject model.Subject) error {
	_, err := subjectCollection.InsertOne(context.TODO(), subject)
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func DeleteSubjectById(id string) (*mongo.DeleteResult, error) {
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	result, errDel := subjectCollection.DeleteOne(context.TODO(), bson.M{"_id": docID})
	if errDel != nil {
		panic(errDel)
	}
	return result, errDel
}

func GetSubjectForUserId(id string) ([]*model.Subject, error) {
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	var subject []*model.Subject
	res, errFind := subjectCollection.Find(context.TODO(), bson.M{"User": docID})
	if errFind != nil {
		log.Fatal(errFind)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for res.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem model.Subject
		err := res.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		subject = append(subject, &elem)
	}

	if err := res.Err(); err != nil {
		log.Fatal(err)
	}
	res.Close(context.TODO())
	return subject, errFind
}

func GetSubjectById(id string) (model.Subject, error) {
	subject := model.Subject{}
	docID, err := primitive.ObjectIDFromHex(id)
	err = subjectCollection.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&subject)
	return subject, err
}
