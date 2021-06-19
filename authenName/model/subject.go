package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subject struct {
	ID           primitive.ObjectID `gorm:"primaryKey" bson:"_id,omitempty"`
	Name         string             `bson:"name,omitempty"`
	Description  string             `bson:"description,omitempty"`
	SubjectID    string             `bson:"subjectId,omitempty"`
	Key          int64              `bson:"key,omitempty"`
	Password     string             `bson:"password,omitempty"`
	Private      string             `bson:"private,omitempty"`
	Status       string             `bson:"status,omitempty"`
	User         primitive.ObjectID `gorm:"foreignKey:id" bson:"User,omitempty"`
	Authen       primitive.ObjectID `gorm:"foreignKey:id" bson:"Authen,omitempty"`
	Upload       primitive.ObjectID `gorm:"foreignKey:id" bson:"Upload,omitempty"`
	CreateAt     time.Time          `bson:"createAt,omitempty"`
	CreateEnd    time.Time          `bson:"createEnd,omitempty"`
	CreateUpdate time.Time          `bson:"createUpdate,omitempty"`
	Score        primitive.ObjectID `gorm:"many2many:Subject_Score" bson:"Score,omitempty"`
	TimeClass    []string           `bson:"timeClass,omitempty"`
}
