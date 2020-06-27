package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UploadInfo struct {
	Upload []Upload `json:"upload"`
}

type Upload struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Path   string             `bson:"path,omitempty"`
	Status string             `bson:"status,omitempty"`
	Name   string             `bson:"name,omitempty"`
}

type UploadReq struct {
	IdPath []string `bson:"IdPath,omitempty"`
}
