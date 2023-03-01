package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"
)

func TestCreateAndStoreUserObject(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	CreateAndStoreUserObject("CreateUserTest")
	client := connectMongoDBAtlas()
	colUser := connectCollection(client, "Users")
	got, _ := colUser.CountDocuments(ctx, bson.D{{"username", "CreateUserTest"}})
	var want int64 = 1

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	colUser.DeleteOne(ctx, bson.D{{"username", "CreateUserTest"}})
	client.Disconnect(ctx)
}
