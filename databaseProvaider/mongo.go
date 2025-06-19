package databaseProvaider

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database
var MongoBucket *gridfs.Bucket

func ConnectMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	MongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	MongoDB = MongoClient.Database("archive")

	MongoBucket, err = gridfs.NewBucket(MongoDB)
	if err != nil {
		log.Fatal("Ошибка создания GridFS bucket:", err)
	}
}
