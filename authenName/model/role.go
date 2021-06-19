package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role struct {
	ID   primitive.ObjectID `gorm:"primaryKey" bson:"_id,omitempty"`
	Name string             `bson:"role,omitempty"`
}
