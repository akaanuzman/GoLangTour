package services

import (
	"chat-app/src/core/config"
	"chat-app/src/core/db"
	"chat-app/src/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// IUserService defines the interface for user services.
type IUserService interface {
	// GetUserByID gets a user by ID.
	GetUserByID(userID string) (*models.User, error)
	// GetUserByEmail gets a user by email.
	GetUserByEmail(email string) (*models.User, error)
	// DeleteUser deletes a user.
	DeleteUser(userID string) error
}

// UserService is a struct that implements the IUserService interface.
type UserService struct {
	userCollection *mongo.Collection
	context        context.Context
	config         *config.Config
}

// NewUserService creates a new instance of UserService with the provided database and configuration.
// Returns an instance of IUserService.
func NewUserService(cfg *config.Config) IUserService {
	db := db.Database{}
	db.ConnectDB(cfg)
	collection := db.GetCollection("users")
	return &UserService{
		userCollection: collection,
		context:        context.TODO(),
		config:         cfg,
	}
}

// GetUserByID gets a user by ID.
func (service *UserService) GetUserByID(userID string) (*models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	var user models.User
	err = service.userCollection.FindOne(service.context, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail gets a user by email.
func (service *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := service.userCollection.FindOne(service.context, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteUser deletes a user.
func (service *UserService) DeleteUser(userID string) error {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	_, err = service.userCollection.DeleteOne(service.context, bson.M{"_id": objectID})
	return err
}
