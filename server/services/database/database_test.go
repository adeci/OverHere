package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"
)

func TestPostUser(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connect
	ConnectMongoDBAtlas()
	PostUser("CreateUserTest")

	// Test
	got, _ := colUsers.CountDocuments(ctx, bson.D{{"username", "CreateUserTest"}})
	var want int64 = 1

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	// Cleanup
	colUsers.DeleteOne(ctx, bson.D{{"username", "CreateUserTest"}})
	db.Disconnect(ctx)
}

func TestPostOHPost(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connect
	ConnectMongoDBAtlas()
	PostOHPost("TEST", "TEST", 21, 21, "tag")

	// Test
	got, _ := colOHPosts.CountDocuments(
		ctx,
		bson.D{
			{"userid", "TEST"},
			{"description", "TEST"},
			{"xcoord", 21},
			{"ycoord", 21}})
	var want int64 = 1

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	// Cleanup
	colOHPosts.DeleteOne(ctx, bson.D{{"userid", "TEST"}})
	db.Disconnect(ctx)
}

func TestPostImage(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connect
	ConnectMongoDBAtlas()
	PostImage("TEST", "TEST", "TEST", 21, 21, "", "")

	// Test
	got, _ := colImages.CountDocuments(
		ctx,
		bson.D{
			{"base64encode", "TEST"},
			{"userid", "TEST"},
			{"ohpostid", "TEST"},
			{"xcoord", 21},
			{"ycoord", 21}})
	var want int64 = 1

	// Cleanup
	colImages.DeleteOne(ctx, bson.D{{"ohpostid", "TEST"}})
	db.Disconnect(ctx)

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPutUser_Username(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connect
	ConnectMongoDBAtlas()
	PostUserTest("TEST", "TEST")

	// Test
	PutUser("TEST", "TESTUPDATE")
	got, _ := GetUser_UserID("TEST")
	want := createUserObject("TESTUPDATE", "TEST")

	// Cleanup
	colUsers.DeleteOne(ctx, bson.D{{"userid", "TEST"}})
	db.Disconnect(ctx)

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGetUser_Username(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connect
	ConnectMongoDBAtlas()

	// Upload test
	PostUserTest("hello7", "hello7")

	// Test
	got, _ := GetUser_Username("hello7")
	want := createUserObject("hello7", "hello7")

	// Cleanup
	colUsers.DeleteOne(ctx, bson.D{{"username", "hello7"}})
	db.Disconnect(ctx)

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGetImage_ImageID(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connect
	ConnectMongoDBAtlas()

	// Upload test
	var x float64 = 0
	PostImageBase("please", "", "", "", x, x, "", "")

	// Test
	got, _ := GetImage_ImageID("please")
	want := createImageObject("please", "", "", "", x, x, "", "")

	// Cleanup
	colImages.DeleteOne(ctx, bson.D{{"imageid", "please"}})
	db.Disconnect(ctx)

	// Assert
	fmt.Println(got)
	fmt.Println(want)
}

func TestDeleteUser_UserID(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connect
	ConnectMongoDBAtlas()

	// Upload test
	PostUserTest("hello7", "hello7")

	// Test
	DeleteUser_UserID("hello7")
	got, _ := colUsers.CountDocuments(ctx, bson.D{{"userid", "hello7"}})
	var want int64 = 0

	db.Disconnect(ctx)

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDeleteOHPost_OHPostID(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connect
	ConnectMongoDBAtlas()

	// Upload test
	PostOHPostBase("TESTTEST", "", "", 0, 0, "tag")

	// Test
	DeleteOHPost_OHPostID("TESTTEST")
	got, _ := colOHPosts.CountDocuments(ctx, bson.D{{"ohpostid", "TESTTEST"}})
	var want int64 = 0

	db.Disconnect(ctx)

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDeleteOHPost_UserID(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connect
	ConnectMongoDBAtlas()

	// Upload test
	PostOHPostBase("", "TESTTEST", "", 0, 0, "tag")
	PostOHPostBase("", "TESTTEST", "", 0, 0, "tag")
	PostOHPostBase("", "TESTTEST", "", 0, 0, "tag")
	PostOHPostBase("", "TESTTEST", "", 0, 0, "tag")

	// Test
	DeleteOHPost_UserID("TESTTEST")
	got, _ := colOHPosts.CountDocuments(ctx, bson.D{{"ohpostid", "TESTTEST"}})
	var want int64 = 0

	db.Disconnect(ctx)

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDeleteImage_ImageID(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connect
	ConnectMongoDBAtlas()

	// Upload test
	PostImageBase("TEST3", "", "", "", 0, 0, "", "")

	// Test
	DeleteImage_ImageID("TEST3")
	got, _ := colImages.CountDocuments(ctx, bson.D{{"imageid", "TEST3"}})
	var want int64 = 0

	db.Disconnect(ctx)

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDeleteImage_UserID(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connect
	ConnectMongoDBAtlas()

	// Upload test
	PostImageBase("", "", "TEST7", "", 0, 0, "", "")
	PostImageBase("", "", "TEST7", "", 0, 0, "", "")
	PostImageBase("", "", "TEST7", "", 0, 0, "", "")

	// Test
	DeleteImage_UserID("TEST7")
	got, _ := colImages.CountDocuments(ctx, bson.D{{"userid", "TEST7"}})
	var want int64 = 0

	db.Disconnect(ctx)

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDeleteImage_OHPostID(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connect
	ConnectMongoDBAtlas()

	// Upload test
	PostImageBase("", "", "", "TEST9", 0, 0, "", "")
	PostImageBase("", "", "", "TEST9", 0, 0, "", "")
	PostImageBase("", "", "", "TEST9", 0, 0, "", "")

	// Test
	DeleteImage_OHPostID("TEST9")
	got, _ := colImages.CountDocuments(ctx, bson.D{{"ohpostid", "TEST9"}})
	var want int64 = 0

	db.Disconnect(ctx)

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestGetOHPost_UserID(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connect
	ConnectMongoDBAtlas()

	//Upload Test
	PostOHPostBase("hi", "wee", "description", 1, 1, "tag")
	PostOHPostBase("bye", "wee", "description", 1, 1, "tag")
	PostOHPostBase("hey", "wee", "description", 1, 1, "tag")

	// Test
	got, _ := GetOHPost_UserID("wee")
	fmt.Println(got[0])
	fmt.Println(got[1])
	fmt.Println(got[2])

	// Clean up
	DeleteOHPost_UserID("wee")

	// Disconnect
	db.Disconnect(ctx)
}

//
////func TestGetOHPost_All(t *testing.T) {
//	// Connect
//	ConnectMongoDBAtlas()
//	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//	colOHPosts := connectCollection(db, "OHPosts")
//
//	// Upload Test
//	PostOHPostBase("wee", "", "", 0, 0)
//	PostOHPostBase("woo", "", "", 0, 0)
//	PostOHPostBase("waa", "", "", 0, 0)
//
//	// Test
//	GetOHPost_All()
//}
