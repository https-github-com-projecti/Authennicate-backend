package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
	"time"
)

type UserInfo struct {
	gorm.Model
	User []User `json:"users"`
}

type User struct {
	ID           primitive.ObjectID `gorm:"primaryKey" bson:"_id,omitempty"`
	CreateAt     time.Time          `bson:"createAt,omitempty"`
	CreateEnd    time.Time          `bson:"createEnd,omitempty"`
	CreateUpdate time.Time          `bson:"createUpdate,omitempty"`
	Username     string             `bson:"username,omitempty"`
	Password     string             `bson:"password,omitempty"`
	Upload       primitive.ObjectID `gorm:"foreignKey:id" bson:"Upload,omitempty"`
	PhoneNumber  string             `bson:"phoneNumber,omitempty"`
	Email        string             `bson:"email,omitempty"`
	Subject      primitive.ObjectID `gorm:"foreignKey:ID;references:name" bson:"Subject,omitempty"`
	Score        primitive.ObjectID `gorm:"many2many:Score_User" bson:"Score,omitempty"`
}

type UserLogin struct {
	gorm.Model
	Username string `bson:"username,omitempty"`
	Password string `bson:"password,omitempty"`
}
