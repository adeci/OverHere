package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type OHPost struct {
	ObjectID primitive.ObjectID `bson:"_id" json:"_id"`
	OHPostId string             `json:"ohpostid,omitempty"`
	UserId   string             `json:"userid,omitempty"`
}
