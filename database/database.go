package database

import (
	"context"
	"golang-sample-crud-service/models"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if client == nil {
		return
	}

	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}

func CreateBook(book models.Book) string {
	client := GetClient()
	bookCollection := GetCollection(client, COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	bookToPost := models.Book{
		Name:  book.Name,
		Price: book.Price,
	}

	result, err := bookCollection.InsertOne(ctx, bookToPost)
	if err != nil {
		return ""
	}

	return result.InsertedID.(primitive.ObjectID).Hex()
}
