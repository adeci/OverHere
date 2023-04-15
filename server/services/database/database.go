package database

import (
	"context"
	"errors"
	"github.com/dchest/uniuri"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var db *mongo.Client

type UserObject struct {
	UserID   string `json: "UserID"`
	Username string `json: "Username"`
}

type OHPostObject struct {
	OHPostID    string  `json: "OHPostID"`
	UserID      string  `json: "UserID"`
	Description string  `json: "Description"`
	XCoord      float64 `json: "XCoord"`
	YCoord      float64 `json: "YCoord"`
}

type ImageObject struct {
	ImageID      string  `json: "ImageID"`
	Base64Encode string  `json: "Base 64 Encode"`
	UserID       string  `json: "UserID"`
	OHPostID     string  `json: "OHPostID"`
	XCoord       float64 `json: "XCoord"`
	YCoord       float64 `json: "YCoord"`
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

func generateOHPostObject(userid string, description string, xcoord float64, ycoord float64) OHPostObject {
	// Generate ohpostid -> Basically userid/generated ohpostid
	ohpostid := userid + "-" + uniuri.New()

	object := OHPostObject{
		OHPostID:    ohpostid,
		UserID:      userid,
		Description: description,
		XCoord:      xcoord,
		YCoord:      ycoord,
	}

	return object
}

func createOHPostObject(ohpostid string, userid string, description string, xcoord float64, ycoord float64) OHPostObject {
	object := OHPostObject{
		OHPostID:    ohpostid,
		UserID:      userid,
		Description: description,
		XCoord:      xcoord,
		YCoord:      ycoord,
	}

	return object
}

func generateImageObject(base64encode string, userid string, ohpostid string, xcoord float64, ycoord float64) ImageObject {
	// Generate imageid -> Basically userid/ohpostid/generated imageid
	imageid := ohpostid + "-" + uniuri.New()

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

func createImageObject(imageid string, base64encode string, userid string, ohpostid string, xcoord float64, ycoord float64) ImageObject {
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

func ConnectMongoDBAtlas() *mongo.Client {
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
	db = client
	return client
}

func connectCollection(client *mongo.Client, collection string) *mongo.Collection {
	col := client.Database("OverHere").Collection(collection)
	return col
}

func PostUser(username string) (UserObject, error) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colUsers := connectCollection(db, "Users")

	// Check If Username Exists Already
	count, _ := colUsers.CountDocuments(ctx, bson.D{{"username", username}})

	if count == 0 {
		// Create User Object
		userObject := generateUserObject(username)
		_, err := colUsers.InsertOne(ctx, userObject)

		return userObject, err
	}

	exists := createUserObject("Username Already Exists", "Username Already Exists")
	return exists, errors.New("Fail Post User: Username Already Exists")
}

func PostUserTest(username string, userid string) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colUsers := connectCollection(db, "Users")

	// Create User Object
	userObject := createUserObject(userid, username)
	colUsers.InsertOne(ctx, userObject)
}

func PostOHPost(userid string, description string, xcoord float64, ycoord float64) (OHPostObject, error) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colOHPosts := connectCollection(db, "OHPosts")

	// Create OHPost Object
	ohpostObject := generateOHPostObject(userid, description, xcoord, ycoord)
	_, err := colOHPosts.InsertOne(ctx, ohpostObject)

	return ohpostObject, err
}

func PostOHPostBase(ohpostid string, userid string, description string, xcoord float64, ycoord float64) (OHPostObject, error) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colOHPosts := connectCollection(db, "OHPosts")

	// Create OHPost Object
	ohpostObject := createOHPostObject(ohpostid, userid, description, xcoord, ycoord)
	_, err := colOHPosts.InsertOne(ctx, ohpostObject)

	return ohpostObject, err
}

func PostImage(base64encode string, userid string, ohpostid string, xcoord float64, ycoord float64) (ImageObject, error) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colImages := connectCollection(db, "Images")

	// Create and Insert User Object
	imageObject := generateImageObject(base64encode, userid, ohpostid, xcoord, ycoord)
	_, err := colImages.InsertOne(ctx, imageObject)

	return imageObject, err
}

func PostImageBase(imageid string, base64encode string, userid string, ohpostid string, xcoord float64, ycoord float64) (ImageObject, error) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colImages := connectCollection(db, "Images")

	// Create and Insert User Object
	imageObject := createImageObject(imageid, base64encode, userid, ohpostid, xcoord, ycoord)
	_, err := colImages.InsertOne(ctx, imageObject)

	return imageObject, err
}

func PutUser(userid string, username string) error {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colUsers := connectCollection(db, "Users")

	// Update
	_, err := colUsers.UpdateOne(ctx, bson.D{{"userid", userid}}, bson.D{{"$set", bson.D{{"username", username}}}})

	return err
}

func PutOHPost(object OHPostObject) {
	// Update
	DeleteOHPost_OHPostID(object.OHPostID)
	PostOHPostBase(object.OHPostID, object.UserID, object.Description, object.XCoord, object.YCoord)
}

func PutImage(object ImageObject) {
	// Update
	DeleteImage_ImageID(object.ImageID)
	PostImageBase(object.ImageID, object.Base64Encode, object.UserID, object.OHPostID, object.XCoord, object.YCoord)
}

func PutOHPost_XCoord(ohpostid string, xcoord float64) error {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colOHPosts := connectCollection(db, "OHPosts")

	// Update
	_, err := colOHPosts.UpdateOne(ctx, bson.D{{"ohpostid", ohpostid}}, bson.D{{"$set", bson.D{{"xcoord", xcoord}}}})

	return err
}

func PutOHPost_YCoord(ohpostid string, ycoord float64) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colOHPosts := connectCollection(db, "OHPosts")

	// Update
	colOHPosts.UpdateOne(ctx, bson.D{{"ohpostid", ohpostid}}, bson.D{{"$set", bson.D{{"ycoord", ycoord}}}})
}

func PutImage_OHPostID(imageid string, ohpostid string) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colOHPosts := connectCollection(db, "OHPosts")

	// Update
	colOHPosts.UpdateOne(ctx, bson.D{{"imageid", imageid}}, bson.D{{"$set", bson.D{{"ohpostid", ohpostid}}}})
}

func PutImage_XCoord(imageid string, xcoord float64) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colOHPosts := connectCollection(db, "OHPosts")

	// Update
	colOHPosts.UpdateOne(ctx, bson.D{{"imageid", imageid}}, bson.D{{"$set", bson.D{{"xcoord", xcoord}}}})
}

func PutImage_YCoord(imageid string, ycoord float64) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colOHPosts := connectCollection(db, "OHPosts")

	// Update
	colOHPosts.UpdateOne(ctx, bson.D{{"imageid", imageid}}, bson.D{{"$set", bson.D{{"ycoord", ycoord}}}})
}

func GetUser_UserID(userid string) (UserObject, error) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colUsers := connectCollection(db, "Users")

	// Check If User Exists
	var want int64 = 1
	got, _ := colUsers.CountDocuments(ctx, bson.D{{"userid", userid}})
	if got == want {
		// Get User
		var user []bson.M
		cursor, err := colUsers.Find(ctx, bson.D{{"userid", userid}})
		cursor.All(ctx, &user)

		return createUserObject(
			user[0]["username"].(string),
			user[0]["userid"].(string)), err
	}

	var object UserObject
	return object, errors.New("GetUser_UserID Fail: Userid Doesn't Exist")
}

func GetUser_Username(username string) (UserObject, error) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colUsers := connectCollection(db, "Users")

	// Check If User Exists
	var want int64 = 1
	got, _ := colUsers.CountDocuments(ctx, bson.D{{"username", username}})

	if got == want {
		// Get User
		var user []bson.M
		cursor, err := colUsers.Find(ctx, bson.D{{"username", username}})
		cursor.All(ctx, &user)

		return createUserObject(
			user[0]["username"].(string),
			user[0]["userid"].(string)), err
	}

	var object UserObject
	return object, errors.New("GetUser_Username Fail: Username Doesn't Exist")
}

func GetUser_All() ([]UserObject, error) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colUsers := connectCollection(db, "Users")

	// Check If User Exists
	var want int64 = 1
	got, _ := colUsers.CountDocuments(ctx, bson.D{})

	if got >= want {
		// Get OHPost
		var users []bson.M
		cursor, err := colUsers.Find(ctx, bson.D{})
		cursor.All(ctx, &users)

		// Create and Fill OHPostObjects Array
		var userObjects []UserObject

		for i := 0; i < len(users); i++ {
			userObjects = append(userObjects, createUserObject(
				users[i]["userid"].(string),
				users[i]["username"].(string)))
		}

		return userObjects, err
	}

	var object []UserObject
	return object, errors.New("GetUser_All Fail: Users Doesn't Exist")
}

func GetOHPost_All() ([]OHPostObject, error) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colOHPosts := connectCollection(db, "OHPosts")

	// Check If OHPost Exists
	var want int64 = 1
	got, _ := colOHPosts.CountDocuments(ctx, bson.D{})

	if got >= want {
		// Get OHPost
		var ohposts []bson.M
		cursor, err := colOHPosts.Find(ctx, bson.D{})
		cursor.All(ctx, &ohposts)

		// Create and Fill OHPostObjects Array
		var ohpostObjects []OHPostObject

		for i := 0; i < len(ohposts); i++ {
			ohpostObjects = append(ohpostObjects, createOHPostObject(
				ohposts[i]["ohpostid"].(string),
				ohposts[i]["userid"].(string),
				ohposts[i]["description"].(string),
				ohposts[i]["xcoord"].(float64),
				ohposts[i]["ycoord"].(float64)))
		}

		return ohpostObjects, err
	}

	var object []OHPostObject
	return object, errors.New("GetOHPostID_All Fail: OHPosts Doesn't Exist")
}

func GetOHPost_OHPostID(ohpostid string) (OHPostObject, error) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colOHPosts := connectCollection(db, "OHPosts")

	// Check If OHPost Exists
	var want int64 = 1
	got, _ := colOHPosts.CountDocuments(ctx, bson.D{{"ohpostid", ohpostid}})

	if got == want {
		// Get OHPost
		var ohpost []bson.M
		cursor, err := colOHPosts.Find(ctx, bson.D{{"ohpostid", ohpostid}})
		cursor.All(ctx, &ohpost)

		return createOHPostObject(
			ohpost[0]["ohpostid"].(string),
			ohpost[0]["userid"].(string),
			ohpost[0]["description"].(string),
			ohpost[0]["xcoord"].(float64),
			ohpost[0]["ycoord"].(float64)), err
	}

	var object OHPostObject
	return object, errors.New("GetOHPostID_OHPostID Fail: OHPostID Doesn't Exist")
}

func GetOHPost_UserID(userid string) ([]OHPostObject, error) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colOHPosts := connectCollection(db, "OHPosts")

	// Check If OHPost Exists
	var want int64 = 1
	got, _ := colOHPosts.CountDocuments(ctx, bson.D{{"userid", userid}})

	if got >= want {
		// Get OHPost
		var ohposts []bson.M
		cursor, err := colOHPosts.Find(ctx, bson.D{{"userid", userid}})
		cursor.All(ctx, &ohposts)

		// Create and Fill OHPostObjects Array
		var ohpostObjects []OHPostObject

		for i := 0; i < len(ohposts); i++ {
			ohpostObjects = append(ohpostObjects, createOHPostObject(
				ohposts[i]["ohpostid"].(string),
				ohposts[i]["userid"].(string),
				ohposts[i]["description"].(string),
				ohposts[i]["xcoord"].(float64),
				ohposts[i]["ycoord"].(float64)))
		}

		return ohpostObjects, err
	}

	var object []OHPostObject
	return object, errors.New("GetOHPostID_OHPostID Fail: UserID Doesn't Exist")
}

func GetImage_ImageID(imageid string) (ImageObject, error) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colImages := connectCollection(db, "Images")

	// Check If User Exists
	var want int64 = 1
	got, _ := colImages.CountDocuments(ctx, bson.D{{"imageid", imageid}})

	if got == want {
		// Get User
		var image []bson.M
		cursor, err := colImages.Find(ctx, bson.D{{"imageid", imageid}})
		cursor.All(ctx, &image)

		return createImageObject(
			image[0]["imageid"].(string),
			image[0]["base64encode"].(string),
			image[0]["userid"].(string),
			image[0]["ohpostid"].(string),
			image[0]["xcoord"].(float64),
			image[0]["ycoord"].(float64)), err
	}

	var object ImageObject
	return object, errors.New("GetImage_ImageID Fail: ImageID Doesn't Exist")
}

func GetImage_UserID(userid string) ([]ImageObject, error) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colImages := connectCollection(db, "Images")

	// Check If Image Exists
	var want int64 = 1
	got, _ := colImages.CountDocuments(ctx, bson.D{{"userid", userid}})

	if got >= want {
		// Get Image
		var images []bson.M
		cursor, err := colImages.Find(ctx, bson.D{{"userid", userid}})
		cursor.All(ctx, &images)

		// Create and Fill ImageObjects Array
		var imageObjects []ImageObject

		for i := 0; i < len(images); i++ {
			imageObjects = append(imageObjects, createImageObject(
				images[i]["imageid"].(string),
				images[i]["base64encode"].(string),
				images[i]["userid"].(string),
				images[i]["ohpostid"].(string),
				images[i]["xcoord"].(float64),
				images[i]["ycoord"].(float64)))
		}

		return imageObjects, err
	}

	var object []ImageObject
	return object, errors.New("GetImage_UserID Fail: UserID Doesn't Exist")

}

func GetImage_OHPostID(ohpostid string) ([]ImageObject, error) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colImages := connectCollection(db, "Images")

	// Check If Image Exists
	var want int64 = 1
	got, _ := colImages.CountDocuments(ctx, bson.D{{"OHPostID", ohpostid}})

	if got >= want {
		// Get Image
		var images []bson.M
		cursor, err := colImages.Find(ctx, bson.D{{"ohpostid", ohpostid}})
		cursor.All(ctx, &images)

		// Create and Fill ImageObjects Array
		var imageObjects []ImageObject

		for i := 0; i < len(images); i++ {
			imageObjects = append(imageObjects, createImageObject(
				images[i]["imageid"].(string),
				images[i]["base64encode"].(string),
				images[i]["userid"].(string),
				images[i]["ohpostid"].(string),
				images[i]["xcoord"].(float64),
				images[i]["ycoord"].(float64)))
		}

		return imageObjects, err
	}

	var object []ImageObject
	return object, errors.New("GetImageOHPostID Fail: OHPostID Doesn't Exist")
}

func GetImage_All() ([]ImageObject, error) {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colImages := connectCollection(db, "Images")

	// Check If Image Exists
	var want int64 = 1
	got, _ := colImages.CountDocuments(ctx, bson.D{})

	if got >= want {
		// Get Image
		var images []bson.M
		cursor, err := colImages.Find(ctx, bson.D{})
		cursor.All(ctx, &images)

		// Create and Fill ImageObjects Array
		var imageObjects []ImageObject

		for i := 0; i < len(images); i++ {
			imageObjects = append(imageObjects, createImageObject(
				images[i]["imageid"].(string),
				images[i]["base64encode"].(string),
				images[i]["userid"].(string),
				images[i]["ohpostid"].(string),
				images[i]["xcoord"].(float64),
				images[i]["ycoord"].(float64)))
		}

		return imageObjects, err
	}

	var object []ImageObject
	return object, errors.New("GetImageOHPostID Fail: Images Doesn't Exist")
}

func DeleteUser_UserID(userid string) error {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colUsers := connectCollection(db, "Users")

	// Delete User
	_, err := colUsers.DeleteOne(ctx, bson.D{{"userid", userid}})

	return err
}

func DeleteUser_Username(username string) error {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colUsers := connectCollection(db, "Users")

	// Delete User
	_, err := colUsers.DeleteOne(ctx, bson.D{{"username", username}})

	return err
}

func DeleteOHPost_OHPostID(ohpostid string) error {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colOHPosts := connectCollection(db, "OHPosts")

	// Delete OHPost
	_, err := colOHPosts.DeleteOne(ctx, bson.D{{"ohpostid", ohpostid}})

	return err
}

func DeleteOHPost_UserID(userid string) error {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colOHPosts := connectCollection(db, "OHPosts")

	// Delete OHPosts
	_, err := colOHPosts.DeleteMany(ctx, bson.D{{"userid", userid}})

	return err
}

func DeleteImage_ImageID(imageid string) error {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colImages := connectCollection(db, "Images")

	// Delete OHPost
	_, err := colImages.DeleteOne(ctx, bson.D{{"imageid", imageid}})

	return err
}

func DeleteImage_UserID(userid string) error {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colImages := connectCollection(db, "Images")

	// Delete OHPosts
	_, err := colImages.DeleteMany(ctx, bson.D{{"userid", userid}})

	return err
}

func DeleteImage_OHPostID(ohpostid string) error {
	// Context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to MongoDB Collections
	colImages := connectCollection(db, "Images")

	// Delete OHPosts
	_, err := colImages.DeleteMany(ctx, bson.D{{"ohpostid", ohpostid}})

	return err
}
