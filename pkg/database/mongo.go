package database

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance *mongo.Client
	clientOnce     sync.Once
)

func GetMongoClient() *mongo.Client {
	clientOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal("MongoDB connection error:", err)
		}

		clientInstance = client
	})
	return clientInstance
}

func GetCollection(database, collection string) *mongo.Collection {
	return GetMongoClient().Database(database).Collection(collection)
}
