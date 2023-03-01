package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"
)

func TestCreateAndStoreUserObject(t *testing.T) {
	// Connect
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	CreateAndStoreUserObject("CreateUserTest")
	client := connectMongoDBAtlas()
	colUsers := connectCollection(client, "Users")

	// Test
	got, _ := colUsers.CountDocuments(ctx, bson.D{{"userid", "CreateUserTest"}, {"username", "CreateUserTest"}})
	var want int64 = 1

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	// Cleanup
	colUsers.DeleteOne(ctx, bson.D{{"username", "CreateUserTest"}})
	client.Disconnect(ctx)
}

func TestCreateAndStoreOHPostObject(t *testing.T) {
	// Connect
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	CreateAndStoreOHPostObject("TEST", "TEST", "TEST")
	client := connectMongoDBAtlas()
	colOHPosts := connectCollection(client, "OHPosts")

	// Test
	got, _ := colOHPosts.CountDocuments(
		ctx,
		bson.D{
			{"ohpostid", "TEST"},
			{"userid", "TEST"},
			{"description", "TEST"}})
	var want int64 = 1

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	// Cleanup
	colOHPosts.DeleteOne(ctx, bson.D{{"ohpostid", "TEST"}})
	client.Disconnect(ctx)
}

func TestCreateAndStoreImageObject(t *testing.T) {
	// Connect
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	CreateAndStoreImageObject("TEST", "TEST", "TEST", "TEST")
	client := connectMongoDBAtlas()
	colImages := connectCollection(client, "Images")

	// Test
	got, _ := colImages.CountDocuments(
		ctx,
		bson.D{
			{"imageid", "TEST"},
			{"base64encode", "TEST"},
			{"userid", "TEST"},
			{"ohpostid", "TEST"}})
	var want int64 = 1

	// Assert
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	// Cleanup
	colImages.DeleteOne(ctx, bson.D{{"imageid", "TEST"}})
	client.Disconnect(ctx)
}
