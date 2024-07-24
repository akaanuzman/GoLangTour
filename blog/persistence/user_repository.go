package persistence

import (
	"blog/common/db"
	"blog/helpers"
	"blog/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepository interface {
	// FindAll retrieves all users from the database.
	FindAll() ([]models.User, error)
	// FindByID retrieves a user by its ID from the database.
	FindByID(id string) (*models.User, error)
	// FindByEmail retrieves a user by its email from the database.
	FindByEmail(email string) (*models.User, error)
	// Create inserts a new user into the database.
	Create(user models.User) (*models.User, error)
	// Update updates an existing user in the database.
	Update(user models.User) (*models.User, error)
	// Delete removes a user from the database.
	Delete(id string) error
}

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository() IUserRepository {
	collection := db.GetCollection("users")
	return &UserRepository{collection: collection}
}

func (userRepository *UserRepository) FindAll() ([]models.User, error) {
	ctx, cancel := helpers.WithTimeout(10 * time.Second)
	defer cancel()

	var users []models.User
	cursor, err := userRepository.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (userRepository *UserRepository) FindByID(id string) (*models.User, error) {
	ctx, cancel := helpers.WithTimeout(10 * time.Second)
	defer cancel()

	var user models.User
	err := userRepository.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepository *UserRepository) FindByEmail(email string) (*models.User, error) {
	ctx, cancel := helpers.WithTimeout(10 * time.Second)
	defer cancel()

	var user models.User
	err := userRepository.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepository *UserRepository) Create(user models.User) (*models.User, error) {
	ctx, cancel := helpers.WithTimeout(10 * time.Second)
	defer cancel()

	_, err := userRepository.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepository *UserRepository) Update(user models.User) (*models.User, error) {
	ctx, cancel := helpers.WithTimeout(10 * time.Second)
	defer cancel()

	_, err := userRepository.collection.ReplaceOne(ctx, bson.M{"_id": user.ID}, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepository *UserRepository) Delete(id string) error {
	ctx, cancel := helpers.WithTimeout(10 * time.Second)
	defer cancel()

	_, err := userRepository.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
