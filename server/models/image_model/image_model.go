package image_model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Image struct {
	ObjectID    primitive.ObjectID `bson:"_id" json:"_id"`
	OHPostId    string             `json:"userid,omitempty"`
	Encoding    string             `json:"encoding" validate:"required"`
	Description string             `json:"description,omitempty"`
}
