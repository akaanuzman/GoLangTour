package db

import (
	"blog/src/config"
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var once sync.Once

func ConnectDB() *mongo.Client {
	once.Do(func() {
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
	})

	return MongoClient
}
