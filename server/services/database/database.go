package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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

func connectMongoDBAtlas() *mongo.Client {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://nicomacias0303:OverNicoHere0303@overhere.i6z1ckb.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func connectCollection(client *mongo.Client, collection string) *mongo.Collection {
	col := client.Database("OverHere").Collection(collection)
	return col
}

// SPRINT 1 DEMO
//func DemoUploadAndRetrieveImage(file string) {
//	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
//
//	col := connectCollection("demo")
//
//	oneDoc := MongoField{
//		FieldBase64Encoding: imageprocessing.DecodePNG(file),
//		FieldLongCoord:      45,
//		FieldLatCoord:       54,
//	}
//
//	result, insertErr := col.InsertOne(ctx, oneDoc)
//	if insertErr != nil {
//		fmt.Println("InsertONE Error: ", insertErr)
//		os.Exit(1)
//	} else {
//
//		newID := result.InsertedID
//		fmt.Println(newID)
//	}
//
//	cursor, err := col.Find(context.TODO(), bson.M{})
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	var images []bson.M
//	if err = cursor.All(ctx, &images); err != nil {
//		fmt.Println(err)
//	}
//
//	for i := 0; i < len(images); i++ {
//		imageprocessing.EncodePNG(images[i])
//	}
//}

func DemoDataStructureOHPostToImages(username string) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colUsers := connectCollection(client, "Users")
	colOHPosts := connectCollection(client, "OHPosts")
	colImages := connectCollection(client, "Images")

	// 1) Create and Insert User Object
	userObject := createUserObject(username)
	colUsers.InsertOne(ctx, userObject)

	// 2) Create and Insert OHPost Object
	ohpostObject := createOHPostObject("ohpostID", username, "description")
	colOHPosts.InsertOne(ctx, ohpostObject)

	// 3) Create and Insert Image Object/s
	imageObject := createImageObject("imageID", "1", username, "ohpostID")
	colImages.InsertOne(ctx, imageObject)
}

func DemoDataStructureImagesToOHPost(username string) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colUsers := connectCollection(client, "Users")
	colOHPosts := connectCollection(client, "OHPosts")
	colImages := connectCollection(client, "Images")

	// 1) Create and Insert User Object
	userObject := createUserObject(username)
	colUsers.InsertOne(ctx, userObject)

	// 2) Create and Insert Image Object/s
	objects := []interface{}{
		createImageObject("a", "1", username, ""),
		createImageObject("b", "2", username, ""),
		createImageObject("c", "3", username, ""),
	}
	colImages.InsertMany(ctx, objects)

	// 3) Wait for OHPost Request, it wants
	// images "a" and "c" in a post.
	// Create and Insert OHPost Object
	ohpostObject := createOHPostObject("test", username, "description")
	colOHPosts.InsertOne(ctx, ohpostObject)

	// 4) Assign images "a" and "c" OHPostID
	// Alongside request, backend receives array
	// of image ids to associate to the OHPost.
	imageIDs := [2]string{"a", "c"}
	for i := 0; i < len(imageIDs); i++ {
		colImages.UpdateOne(
			ctx,
			bson.D{{"imageid", imageIDs[i]}},
			bson.D{{"$set", bson.D{{"ohpostid", "test"}}}})
	}
}
