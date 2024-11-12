package db

import (
	"chat-app/src/core/config"
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database struct represents a connection to the MongoDB database.
// It ensures that the connection is established only once.
type Database struct {
	Client *mongo.Client
	once   sync.Once
}

// ConnectDB establishes a connection to the MongoDB database.
// It uses a sync.Once to ensure the connection is created only once.
// Returns a pointer to the mongo.Client.
func (database *Database) ConnectDB(cfg *config.Config) *mongo.Client {
	database.once.Do(func() {

		// Create a context with a timeout for the connection attempt
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Create a mongo.Client options struct with the URI
		clientOptions := options.Client().ApplyURI(cfg.MongoURI)

		// Connect to the MongoDB database
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			panic(err) // Replace with proper error handling
		}

		// Ping the primary to verify the connection
		if err := client.Ping(ctx, nil); err != nil {
			panic(err) // Replace with proper error handling
		}

		log.Println("Connected to MongoDB!")

		// Store the client in the Database struct
		database.Client = client
	})

	return database.Client
}

// GetCollection returns a pointer to a mongo.Collection for the given collection name.
// Assumes the database name is "chat-app".
func (database *Database) GetCollection(collectionName string) *mongo.Collection {
	// Get the collection from the "chat-app" database
	return database.Client.Database("chat-app").Collection(collectionName)
}
