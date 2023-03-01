package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Image struct {
	ObjectID primitive.ObjectID `bson:"_id" json:"_id"`
	ImageID  string             `json:"imageid,omitempty"`
	OHPostID string             `json:"ohpostid,omitempty"`
	Encoding string             `json:"encoding" validate:"required"`
}
