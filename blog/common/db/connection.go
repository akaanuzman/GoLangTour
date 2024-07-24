package db

import (
	"blog/common/config"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectDB() {
	cfg := config.Config{}
	cfg.LoadConfig()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	// Ping the primary to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	log.Println("Connected to MongoDB!")

	MongoClient = client
}

func GetCollection(collectionName string) *mongo.Collection {
	return MongoClient.Database("blogapp").Collection(collectionName)
}
