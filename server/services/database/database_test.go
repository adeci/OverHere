package database

import (
	"context"
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
	colOHPosts.DeleteOne(ctx, bson.D{{"ohpostid", "TEST"}})
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

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	// Cleanup
	colImages.DeleteOne(ctx, bson.D{{"imageid", "TEST"}})
	client.Disconnect(ctx)
}

func TestGetUserObject(t *testing.T) {
	// Connect
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := connectMongoDBAtlas()
	colUsers := connectCollection(client, "Users")

	// Upload test
	PostUser("hello7")

	// Test
	got := GetUser_Username("hello7")
	want := createUserObject("hello7", "hi")

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	// Cleanup
	colUsers.DeleteOne(ctx, bson.D{{"userid", "hello7"}})
}
