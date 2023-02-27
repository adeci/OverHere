package database

import (
	"OverHere/server/services/imageprocessing"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

type MongoField struct {
	FieldBase64Encoding string `json: "Field Base 64 Encoding"`
	FieldLongCoord      int    `json: "Field LongCord"`
	FieldLatCoord       int    `json: "Field LatCord"`
}

type UserObject struct {
	UserID   string `json: "UserID"`
	Username string `json: "Username"`
}

type OHPostObject struct {
	OHPostID    string `json: "OHPostID"`
	UserID      string `json: "UserID"`
	Description string `json: "Description"`
}

type ImageObject struct {
	ImageID      string `json: "ImageID"`
	Base64Encode string `json: "Base 64 Encode"`
	UserID       string `json: "UserID"`
	OHPostID     string `json: "OHPostID"`
}

// SPRINT 1 DEMO
func DemoUploadAndRetrieveImage(file string) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	col := connectDemoMongoDB()

	oneDoc := MongoField{
		FieldBase64Encoding: imageprocessing.DecodePNG(file),
		FieldLongCoord:      45,
		FieldLatCoord:       54,
	}

	result, insertErr := col.InsertOne(ctx, oneDoc)
	if insertErr != nil {
		fmt.Println("InsertONE Error: ", insertErr)
		os.Exit(1)
	} else {

		newID := result.InsertedID
		fmt.Println(newID)
	}

	cursor, err := col.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	var images []bson.M
	if err = cursor.All(ctx, &images); err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(images); i++ {
		imageprocessing.EncodePNG(images[i])
	}
}

func connectDemoMongoDB() *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() ERROR: ", err)
		os.Exit(1)
	}

	col := client.Database("OverHere").Collection("demo")

	return col
}

func createUserObject(username string) UserObject {
	object := UserObject{
		UserID:   username,
		Username: username,
	}

	return object
}

func createOHPostObject(ohpostid string, userid string, description string) OHPostObject {
	object := OHPostObject{
		OHPostID:    ohpostid,
		UserID:      userid,
		Description: description,
	}

	return object
}

func createImageObject(imageid string, base64encode string, userid string, ohpostid string) ImageObject {
	object := ImageObject{
		ImageID:      imageid,
		Base64Encode: base64encode,
		UserID:       userid,
		OHPostID:     ohpostid,
	}

	return object
}
