package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func main() {
	var conn *mongo.Client
	conn, err := db.ConnectDB()
	defer conn.Disconnect(context.Background())

	c := conn.Database("app").Collection("posts")

	repo := NewMongoDBPostRepository(c)

	repo.Create(context.TODO(), Post{
		ChatID:    0,
		Locations: Location{},
		Status:    "",
	})
	repo.GetAll(context.TODO())
}
