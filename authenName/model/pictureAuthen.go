package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type PictureAuthen struct {
	ID           primitive.ObjectID `gorm:"primaryKey" bson:"_id,omitempty"`
	CreateAt     time.Time          `bson:"createAt,omitempty"`
	CreateEnd    time.Time          `bson:"createEnd,omitempty"`
	CreateUpdate time.Time          `bson:"createUpdate,omitempty"`
	User         primitive.ObjectID `gorm:"foreignKey:id" bson:"User,omitempty"`
	Authen       primitive.ObjectID `gorm:"foreignKey:id" bson:"Authen,omitempty"`
}
