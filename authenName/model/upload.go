package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
	"time"
)

type UploadInfo struct {
	gorm.Model
	Upload []Upload `json:"upload"`
}

type Upload struct {
	ID           primitive.ObjectID `gorm:"primaryKey" bson:"_id,omitempty"`
	Path         string             `bson:"path,omitempty"`
	Status       string             `bson:"status,omitempty"`
	Name         string             `bson:"name,omitempty"`
	CreateAt     time.Time          `bson:"createAt,omitempty"`
	CreateEnd    time.Time          `bson:"createEnd,omitempty"`
	CreateUpdate time.Time          `bson:"createUpdate,omitempty"`
	Authen       primitive.ObjectID `gorm:"many2many:Upload_Authen" bson:"Authen,omitempty"`
}

type UploadReq struct {
	gorm.Model
	IdPath []string `bson:"IdPath,omitempty"`
}
