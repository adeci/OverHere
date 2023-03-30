package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"
)

func TestPostUser(t *testing.T) {
	// Connect
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	PostUser("CreateUserTest")
	client := connectMongoDBAtlas()
	colUsers := connectCollection(client, "Users")

	// Test
	got, _ := colUsers.CountDocuments(ctx, bson.D{{"username", "CreateUserTest"}})
	var want int64 = 1

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	// Cleanup
	colUsers.DeleteOne(ctx, bson.D{{"username", "CreateUserTest"}})
	client.Disconnect(ctx)
}

func TestPostOHPost(t *testing.T) {
	// Connect
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	PostOHPost("TEST", "TEST", 21, 21)
	client := connectMongoDBAtlas()
	colOHPosts := connectCollection(client, "OHPosts")

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
	client.Disconnect(ctx)
}

func TestPostImage(t *testing.T) {
	// Connect
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	PostImage("TEST", "TEST", "TEST", 21, 21)
	client := connectMongoDBAtlas()
	colImages := connectCollection(client, "Images")

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
	client.Disconnect(ctx)

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPutUser_Username(t *testing.T) {
	// Connect
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	PostUserTest("TEST", "TEST")
	client := connectMongoDBAtlas()
	colUsers := connectCollection(client, "Users")

	// Test
	PutUser("TEST", "TESTUPDATE")
	got, _ := GetUser_UserID("TEST")
	want := createUserObject("TESTUPDATE", "TEST")

	// Cleanup
	colUsers.DeleteOne(ctx, bson.D{{"userid", "TEST"}})
	client.Disconnect(ctx)

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGetUser_Username(t *testing.T) {
	// Connect
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := connectMongoDBAtlas()
	colUsers := connectCollection(client, "Users")

	// Upload test
	PostUserTest("hello7", "hello7")

	// Test
	got, _ := GetUser_Username("hello7")
	want := createUserObject("hello7", "hello7")

	// Cleanup
	colUsers.DeleteOne(ctx, bson.D{{"username", "hello7"}})
	client.Disconnect(ctx)

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGetImage_ImageID(t *testing.T) {
	// Connect
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := connectMongoDBAtlas()
	colImages := connectCollection(client, "Images")

	// Upload test
	var x float64 = 0
	PostImageBase("please", "", "", "", x, x)

	// Test
	got, _ := GetImage_ImageID("please")
	want := createImageObject("please", "", "", "", x, x)

	// Cleanup
	colImages.DeleteOne(ctx, bson.D{{"imageid", "please"}})
	client.Disconnect(ctx)

	// Assert
	fmt.Println(got)
	fmt.Println(want)
}

func TestDeleteUser_UserID(t *testing.T) {
	// Connect
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := connectMongoDBAtlas()
	colUsers := connectCollection(client, "Users")

	// Upload test
	PostUserTest("hello7", "hello7")

	// Test
	DeleteUser_UserID("hello7")
	got, _ := colUsers.CountDocuments(ctx, bson.D{{"userid", "hello7"}})
	var want int64 = 0

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDeleteOHPost_OHPostID(t *testing.T) {
	// Connect
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := connectMongoDBAtlas()
	colOHPosts := connectCollection(client, "OHPosts")

	// Upload test
	PostOHPostBase("TESTTEST", "", "", 0, 0)

	// Test
	DeleteOHPost_OHPostID("TESTTEST")
	got, _ := colOHPosts.CountDocuments(ctx, bson.D{{"ohpostid", "TESTTEST"}})
	var want int64 = 0

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDeleteOHPost_UserID(t *testing.T) {
	// Connect
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := connectMongoDBAtlas()
	colOHPosts := connectCollection(client, "OHPosts")

	// Upload test
	PostOHPostBase("", "TESTTEST", "", 0, 0)
	PostOHPostBase("", "TESTTEST", "", 0, 0)
	PostOHPostBase("", "TESTTEST", "", 0, 0)
	PostOHPostBase("", "TESTTEST", "", 0, 0)

	// Test
	DeleteOHPost_UserID("TESTTEST")
	got, _ := colOHPosts.CountDocuments(ctx, bson.D{{"ohpostid", "TESTTEST"}})
	var want int64 = 0

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDeleteImage_ImageID(t *testing.T) {
	// Connect
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := connectMongoDBAtlas()
	colImages := connectCollection(client, "Images")

	// Upload test
	PostImageBase("TEST3", "", "", "", 0, 0)

	// Test
	DeleteImage_ImageID("TEST3")
	got, _ := colImages.CountDocuments(ctx, bson.D{{"imageid", "TEST3"}})
	var want int64 = 0

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDeleteImage_UserID(t *testing.T) {
	// Connect
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := connectMongoDBAtlas()
	colImages := connectCollection(client, "Images")

	// Upload test
	PostImageBase("", "", "TEST7", "", 0, 0)
	PostImageBase("", "", "TEST7", "", 0, 0)
	PostImageBase("", "", "TEST7", "", 0, 0)

	// Test
	DeleteImage_UserID("TEST7")
	got, _ := colImages.CountDocuments(ctx, bson.D{{"userid", "TEST7"}})
	var want int64 = 0

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDeleteImage_OHPostD(t *testing.T) {
	// Connect
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := connectMongoDBAtlas()
	colImages := connectCollection(client, "Images")

	// Upload test
	PostImageBase("", "", "", "TEST9", 0, 0)
	PostImageBase("", "", "", "TEST9", 0, 0)
	PostImageBase("", "", "", "TEST9", 0, 0)

	// Test
	DeleteImage_OHPostID("TEST9")
	got, _ := colImages.CountDocuments(ctx, bson.D{{"ohpostid", "TEST9"}})
	var want int64 = 0

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
