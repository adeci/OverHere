package user_model

import "go.mongodb.org/mongo-driver/bson/primitive"

//We added omitempty and validate:"required" to the struct tag to tell Gin-gonic to ignore empty fields and make the field required, respectively.
type User struct {
	ObjectID primitive.ObjectID `bson:"_id" json:"_id"`
	UserId   string             `json:"userid" validate:"required"`
	Username string             `json:"username,omitempty" validate:"required"`
}
