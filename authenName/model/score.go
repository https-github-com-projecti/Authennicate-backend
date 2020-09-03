package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Score struct {
	ID           primitive.ObjectID `gorm:"primaryKey" bson:"_id,omitempty"`
	Authen       []Authen           `gorm:"foreignKey:id" bson:"Authen,omitempty"`
	CreateAt     time.Time          `bson:"createAt,omitempty"`
	CreateEnd    time.Time          `bson:"createEnd,omitempty"`
	CreateUpdate time.Time          `bson:"createUpdate,omitempty"`
	User         primitive.ObjectID `gorm:"many2many:Score_User" bson:"User,omitempty"`
	Subject      primitive.ObjectID `gorm:"many2many:Subject_Score" bson:"Subject,omitempty"`
}
