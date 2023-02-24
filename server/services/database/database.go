package database

import (
	"OverHere/server/controllers/imageprocessing"
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoField struct {
	FieldBase64Encoding string `json: "Field Base 64 Encoding"`
	FieldLongCoord      int    `json: "Field LongCord"`
	FieldLatCoord       int    `json: "Field LatCord"`
}

func DemoUploadAndRetrieveImage(file string) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() ERROR: ", err)
		os.Exit(1)
	}

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	col := client.Database("OverHere").Collection("demo")

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
