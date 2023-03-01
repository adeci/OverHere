package models

//ObjectID primitive.ObjectID `bson:"_id" json:"_id"`

//We added omitempty and validate:"required" to the struct tag to tell Gin-gonic to ignore empty fields and make the field required, respectively.
type User struct {
	UserID   string `json:"userid,omitempty"`
	Username string `json:"username,omitempty" validate:"required"`
}
