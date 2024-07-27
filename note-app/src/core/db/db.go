package db

import (
	"context"
	"log"
	"note-app/src/core/config"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
	once   sync.Once
}

func (database *Database) ConnectDB() *mongo.Client {
	database.once.Do(func() {
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

		database.Client = client
	})

	return database.Client
}

func (database *Database) GetCollection(collectionName string) *mongo.Collection {
	return database.Client.Database("note-app").Collection(collectionName)
}
