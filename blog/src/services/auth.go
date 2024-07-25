package services

import (
	"blog/src/db"
	"blog/src/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IAuthService interface {
	Register(user *models.User) error
	Login(user *models.User) (string, *models.User, error)
}

type UserAuthService struct {
	userCollection *mongo.Collection
}

func NewUserService() IAuthService {
	client := db.ConnectDB()
	return &UserAuthService{
		userCollection: client.Database("blog").Collection("users"),
	}
}

func (userService *UserAuthService) Register(user *models.User) error {
	if ok, err := validateEmailAndPassword(user); ok {
		return err
	}

	var existingUser models.User

	err := userService.userCollection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return errors.New("email is already in use")
	}

	if err := user.HashPassword(); err != nil {
		return errors.New("error hashing password")
	}

	user.Id = primitive.NewObjectID()
	_, err = userService.userCollection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}

func (userService *UserAuthService) Login(user *models.User) (string, *models.User, error) {
	var dbUser models.User
	err := userService.userCollection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&dbUser)
	if err != nil {
		return "", nil, errors.New("wrong password or email")
	}

	if err := dbUser.ComparePassword(user.Password); err != nil {
		return "", nil, errors.New("wrong password or email")
	}

	token, err := dbUser.GenerateJWT()
	if err != nil {
		return "", nil, errors.New("error generating JWT")
	}

	return token, &dbUser, nil
}

func validateEmailAndPassword(user *models.User) (bool, error) {
	if user.Email == "" || user.Password == "" {
		return true, errors.New("email and password are required")
	}

	if len(user.Password) < 6 {
		return true, errors.New("password must be at least 6 characters long")
	}

	if err := models.ValidateEmail(user.Email); err != nil {
		return true, err
	}
	return false, nil
}
