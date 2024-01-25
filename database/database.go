package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var bookCollection *mongo.Collection

var COLLECTION = "Books"

func GetClient() *mongo.Client {
	uri := os.Getenv("DATABASE_URL")

	if client != nil {
		return client
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatalln(err)
	}

	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	if bookCollection != nil {
		return bookCollection
	}

	bookCollection := client.Database("admin").Collection(collectionName)

	return bookCollection
}
