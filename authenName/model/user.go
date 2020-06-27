package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserInfo struct {
	User []User `json:"users"`
}

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	CreateAt    time.Time          `bson:"create,omitempty"`
	Username    string             `bson:"username,omitempty"`
	Password    string             `bson:"password,omitempty"`
	Upload      primitive.ObjectID `bson:"Upload,omitempty"`
	PhoneNumber string             `bson:"phoneNumber,omitempty"`
	Email       string             `bson:"email,omitempty"`
}

type UserLogin struct {
	Username string `bson:"username,omitempty"`
	Password string `bson:"password,omitempty"`
}
