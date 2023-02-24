package user_model

import "go.mongodb.org/mongo-driver/bson/primitive"

//We added omitempty and validate:"required" to the struct tag to tell Gin-gonic to ignore empty fields and make the field required, respectively.
type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Location string             `json:"location,omitempty" validate:"required"`
	Title    string             `json:"title,omitempty" validate:"required"`
}
