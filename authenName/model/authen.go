package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Authen struct {
	ID           primitive.ObjectID `gorm:"primaryKey" bson:"_id,omitempty"`
	CreateAt     time.Time          `bson:"createAt,omitempty"`
	CreateEnd    time.Time          `bson:"createEnd,omitempty"`
	CreateUpdate time.Time          `bson:"createUpdate,omitempty"`
	Status       bool               `bson:"status,omitempty"`
	Subject      primitive.ObjectID `gorm:"foreignKey:id" bson:"Subject,omitempty"`
	Upload       primitive.ObjectID `gorm:"many2many:Upload_Authen" bson:"Upload,omitempty"`
}
