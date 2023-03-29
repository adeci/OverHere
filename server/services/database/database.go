package database

import (
	"context"
	"github.com/dchest/uniuri"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type UserObject struct {
	UserID   string `json: "UserID"`
	Username string `json: "Username"`
}

type OHPostObject struct {
	OHPostID    string  `json: "OHPostID"`
	UserID      string  `json: "UserID"`
	Description string  `json: "Description"`
	XCoord      float32 `json: "XCoord"`
	YCoord      float32 `json: "YCoord"`
}

type ImageObject struct {
	ImageID      string  `json: "ImageID"`
	Base64Encode string  `json: "Base 64 Encode"`
	UserID       string  `json: "UserID"`
	OHPostID     string  `json: "OHPostID"`
	XCoord       float32 `json: "XCoord"`
	YCoord       float32 `json: "YCoord"`
}

func generateUserObject(username string) UserObject {
	// Generate userid
	userid := uniuri.New()

	// Create User Object
	object := UserObject{
		UserID:   userid,
		Username: username,
	}

	return object
}

func createUserObject(username string, userid string) UserObject {
	// Create User Object
	object := UserObject{
		UserID:   userid,
		Username: username,
	}

	return object
}

func generateOHPostObject(userid string, description string, xcoord float32, ycoord float32) OHPostObject {
	// Generate ohpostid -> Basically userid/generated ohpostid
	ohpostid := userid + "/" + uniuri.New()

	object := OHPostObject{
		OHPostID:    ohpostid,
		UserID:      userid,
		Description: description,
		XCoord:      xcoord,
		YCoord:      ycoord,
	}

	return object
}

func createOHPostObject(ohpostid string, userid string, description string, xcoord float32, ycoord float32) OHPostObject {
	object := OHPostObject{
		OHPostID:    ohpostid,
		UserID:      userid,
		Description: description,
		XCoord:      xcoord,
		YCoord:      ycoord,
	}

	return object
}

func generateImageObject(base64encode string, userid string, ohpostid string, xcoord float32, ycoord float32) ImageObject {
	// Generate imageid -> Basically userid/ohpostid/generated imageid
	imageid := userid + "/" + uniuri.New()

	object := ImageObject{
		ImageID:      imageid,
		Base64Encode: base64encode,
		UserID:       userid,
		OHPostID:     ohpostid,
		XCoord:       xcoord,
		YCoord:       ycoord,
	}

	return object
}

func createImageObject(imageid string, base64encode string, userid string, ohpostid string, xcoord float32, ycoord float32) ImageObject {
	object := ImageObject{
		ImageID:      imageid,
		Base64Encode: base64encode,
		UserID:       userid,
		OHPostID:     ohpostid,
		XCoord:       xcoord,
		YCoord:       ycoord,
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

func PostUser(username string) UserObject {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colUsers := connectCollection(client, "Users")

	// Check If Username Exists Already
	count, _ := colUsers.CountDocuments(ctx, bson.D{{"userid", username}})

	if count == 0 {
		// Create User Object
		userObject := generateUserObject(username)
		colUsers.InsertOne(ctx, userObject)

		// Disconnect
		client.Disconnect(ctx)

		return userObject
	}

	// Disconnect
	client.Disconnect(ctx)

	// TODO: USERNAME EXISTS HANDLING
	exists := generateUserObject("Username Already Exists")
	return exists
}

func PostOHPost(userid string, description string, xcoord float32, ycoord float32) OHPostObject {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colOHPosts := connectCollection(client, "OHPosts")

	// Create OHPost Object
	ohpostObject := generateOHPostObject(userid, description, xcoord, ycoord)
	colOHPosts.InsertOne(ctx, ohpostObject)

	// Disconnect
	client.Disconnect(ctx)

	return ohpostObject

}

func PostImage(base64encode string, userid string, ohpostid string, xcoord float32, ycoord float32) ImageObject {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colImages := connectCollection(client, "Images")

	// Create and Insert User Object
	imageObject := generateImageObject(base64encode, userid, ohpostid, xcoord, ycoord)
	colImages.InsertOne(ctx, imageObject)

	// Disconnect
	client.Disconnect(ctx)

	return imageObject
}

func PutUser_Username(userid string, username string) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colUsers := connectCollection(client, "Users")

	// Update
	colUsers.UpdateOne(ctx, bson.D{{"userid", userid}}, bson.D{{"$set", bson.D{{"username", username}}}})

	// Disconnect
	client.Disconnect(ctx)
}

func PutOHPost_Description(ohpostid string, description string) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colOHPosts := connectCollection(client, "OHPosts")

	// Update
	colOHPosts.UpdateOne(ctx, bson.D{{"ohpostid", ohpostid}}, bson.D{{"$set", bson.D{{"description", description}}}})

	// Disconnect
	client.Disconnect(ctx)
}

func PutOHPost_XCoord(ohpostid string, xcoord float32) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colOHPosts := connectCollection(client, "OHPosts")

	// Update
	colOHPosts.UpdateOne(ctx, bson.D{{"ohpostid", ohpostid}}, bson.D{{"$set", bson.D{{"xcoord", xcoord}}}})

	// Disconnect
	client.Disconnect(ctx)
}

func PutOHPost_YCoord(ohpostid string, ycoord float32) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colOHPosts := connectCollection(client, "OHPosts")

	// Update
	colOHPosts.UpdateOne(ctx, bson.D{{"ohpostid", ohpostid}}, bson.D{{"$set", bson.D{{"ycoord", ycoord}}}})

	// Disconnect
	client.Disconnect(ctx)
}

func PutImage_OHPostID(imageid string, ohpostid string) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colOHPosts := connectCollection(client, "OHPosts")

	// Update
	colOHPosts.UpdateOne(ctx, bson.D{{"imageid", imageid}}, bson.D{{"$set", bson.D{{"ohpostid", ohpostid}}}})

	// Disconnect
	client.Disconnect(ctx)
}

func PutImage_XCoord(imageid string, xcoord float32) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colOHPosts := connectCollection(client, "OHPosts")

	// Update
	colOHPosts.UpdateOne(ctx, bson.D{{"imageid", imageid}}, bson.D{{"$set", bson.D{{"xcoord", xcoord}}}})

	// Disconnect
	client.Disconnect(ctx)
}

func PutImage_YCoord(imageid string, ycoord float32) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colOHPosts := connectCollection(client, "OHPosts")

	// Update
	colOHPosts.UpdateOne(ctx, bson.D{{"imageid", imageid}}, bson.D{{"$set", bson.D{{"ycoord", ycoord}}}})

	// Disconnect
	client.Disconnect(ctx)
}

func GetUser_UserID(userid string) UserObject {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colUsers := connectCollection(client, "Users")

	// Get User
	var user []bson.M
	cursor, _ := colUsers.Find(ctx, bson.D{{"userid", userid}})
	cursor.All(ctx, &user)

	// Disconnect
	client.Disconnect(ctx)

	return createUserObject(
		user[0]["username"].(string),
		user[0]["userid"].(string))

}

func GetUser_Username(username string) UserObject {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colUsers := connectCollection(client, "Users")

	// Get User
	var user []bson.M
	cursor, _ := colUsers.Find(ctx, bson.D{{"username", username}})
	cursor.All(ctx, &user)

	// Disconnect
	client.Disconnect(ctx)

	return createUserObject(
		user[0]["username"].(string),
		user[0]["userid"].(string))
}

func GetOHPost_OHPostID(ohpostid string) OHPostObject {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colOHPosts := connectCollection(client, "OHPosts")

	// Get OHPost
	var ohpost []bson.M
	cursor, _ := colOHPosts.Find(ctx, bson.D{{"ohpostid", ohpostid}})
	cursor.All(ctx, &ohpost)

	// Disconnect
	client.Disconnect(ctx)

	return createOHPostObject(
		ohpost[0]["ohpostid"].(string),
		ohpost[0]["userid"].(string),
		ohpost[0]["description"].(string),
		ohpost[0]["xcoord"].(float32),
		ohpost[0]["ycoord"].(float32))
}
func GetOHPost_UserID(userid string) []OHPostObject {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colOHPosts := connectCollection(client, "OHPosts")

	// Get OHPost
	var ohposts []bson.M
	cursor, _ := colOHPosts.Find(ctx, bson.D{{"userid", userid}})
	cursor.All(ctx, &ohposts)

	// Create and Fill OHPostObjects Array
	var ohpostObjects []OHPostObject

	for i := 0; i < len(ohposts); i++ {
		ohpostObjects = append(ohpostObjects, createOHPostObject(
			ohposts[i]["ohpostid"].(string),
			ohposts[i]["userid"].(string),
			ohposts[i]["description"].(string),
			ohposts[i]["xcoord"].(float32),
			ohposts[i]["ycoord"].(float32)))
	}

	// Disconnect
	client.Disconnect(ctx)

	return ohpostObjects
}

func GetImage_ImageID(imageid string) ImageObject {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colOHPosts := connectCollection(client, "Images")

	// Get Image
	var image []bson.M
	cursor, _ := colOHPosts.Find(ctx, bson.D{{"imageid", imageid}})
	cursor.All(ctx, &image)

	// Disconnect
	client.Disconnect(ctx)

	return createImageObject(
		image[0]["imageid"].(string),
		image[0]["base64encode"].(string),
		image[0]["userid"].(string),
		image[0]["ohpostid"].(string),
		image[0]["xcoord"].(float32),
		image[0]["ycoord"].(float32))
}

func GetImage_UserID(userid string) []ImageObject {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colImages := connectCollection(client, "Images")

	// Get OHPost
	var images []bson.M
	cursor, _ := colImages.Find(ctx, bson.D{{"userid", userid}})
	cursor.All(ctx, &images)

	// Create and Fill ImageObjects Array
	var imageObjects []ImageObject

	for i := 0; i < len(images); i++ {
		imageObjects = append(imageObjects, createImageObject(
			images[i]["imageid"].(string),
			images[i]["base64encode"].(string),
			images[i]["userid"].(string),
			images[i]["ohpostid"].(string),
			images[i]["xcoord"].(float32),
			images[i]["ycoord"].(float32)))
	}

	// Disconnect
	client.Disconnect(ctx)

	return imageObjects
}

func GetImage_OHPostID(ohpostid string) []ImageObject {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colImages := connectCollection(client, "Images")

	// Get OHPost
	var images []bson.M
	cursor, _ := colImages.Find(ctx, bson.D{{"ohpostid", ohpostid}})
	cursor.All(ctx, &images)

	// Create and Fill ImageObjects Array
	var imageObjects []ImageObject

	for i := 0; i < len(images); i++ {
		imageObjects = append(imageObjects, createImageObject(
			images[i]["imageid"].(string),
			images[i]["base64encode"].(string),
			images[i]["userid"].(string),
			images[i]["ohpostid"].(string),
			images[i]["xcoord"].(float32),
			images[i]["ycoord"].(float32)))
	}

	// Disconnect
	client.Disconnect(ctx)

	return imageObjects
}

func DeleteUser_UserID(userid string) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colUsers := connectCollection(client, "Users")

	// Delete User
	colUsers.DeleteOne(ctx, bson.D{{"userid", userid}})

	// Disconnect
	client.Disconnect(ctx)
}

func DeleteUser_Username(username string) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colUsers := connectCollection(client, "Users")

	// Delete User
	colUsers.DeleteOne(ctx, bson.D{{"username", username}})

	// Disconnect
	client.Disconnect(ctx)
}

func DeleteOHPost_OHPostID(ohpostid string) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colOHPosts := connectCollection(client, "OHPosts")

	// Delete OHPost
	colOHPosts.DeleteOne(ctx, bson.D{{"ohpostid", ohpostid}})

	// Disconnect
	client.Disconnect(ctx)
}

func DeleteOHPost_UserID(userid string) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colOHPosts := connectCollection(client, "OHPosts")

	// Delete OHPosts
	colOHPosts.DeleteMany(ctx, bson.D{{"userid", userid}})

	// Disconnect
	client.Disconnect(ctx)
}

func DeleteImage_ImageID(imageid string) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colImages := connectCollection(client, "Images")

	// Delete OHPost
	colImages.DeleteOne(ctx, bson.D{{"imageid", imageid}})

	// Disconnect
	client.Disconnect(ctx)
}

func DeleteImage_UserID(userid string) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colImages := connectCollection(client, "Images")

	// Delete OHPosts
	colImages.DeleteMany(ctx, bson.D{{"userid", userid}})

	// Disconnect
	client.Disconnect(ctx)
}

func DeleteImage_OHPostID(ohpostid string) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Atlas
	client := connectMongoDBAtlas()

	// Connecting to MongoDB Collections
	colImages := connectCollection(client, "Images")

	// Delete OHPosts
	colImages.DeleteMany(ctx, bson.D{{"ohpostid", ohpostid}})

	// Disconnect
	client.Disconnect(ctx)
}
