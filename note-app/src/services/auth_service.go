package services

import (
	"context"
	"errors"
	"note-app/src/core/db"
	"note-app/src/models"
	"note-app/src/models/requests"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IAuthService interface {
	Register(user *models.User) error
	Login(user *models.User) (string, *models.User, error)
	ChangePassword(userID primitive.ObjectID, req *requests.PasswordChangeRequest) error
	ForgotPassword(req *requests.ForgotPasswordRequest) error
}

type UserAuthService struct {
	userCollection *mongo.Collection
	context        context.Context
}

func NewUserService() IAuthService {
	db := db.Database{}
	db.ConnectDB()
	collection := db.GetCollection("users")
	return &UserAuthService{
		userCollection: collection,
		context:        context.Background(),
	}
}

func (userService *UserAuthService) Register(user *models.User) error {
	if ok, err := validateEmailAndPassword(user); ok {
		return err
	}

	var existingUser models.User

	err := userService.userCollection.FindOne(userService.context, bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return errors.New("email is already in use")
	}

	if err := user.HashPassword(); err != nil {
		return errors.New("error hashing password")
	}

	user.Id = primitive.NewObjectID()
	_, err = userService.userCollection.InsertOne(userService.context, user)
	if err != nil {
		return err
	}

	return nil
}

func (userService *UserAuthService) Login(user *models.User) (string, *models.User, error) {
	var dbUser models.User
	err := userService.userCollection.FindOne(userService.context, bson.M{"email": user.Email}).Decode(&dbUser)
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

func (userService *UserAuthService) ChangePassword(userID primitive.ObjectID, req *requests.PasswordChangeRequest) error {
	var dbUser models.User
	err := userService.userCollection.FindOne(userService.context, bson.M{"_id": userID}).Decode(&dbUser)
	if err != nil {
		return errors.New("user not found")
	}

	// Check if reset token is set and if it's expired
	if dbUser.ResetToken != "" && dbUser.ResetTokenExpiry.After(time.Now()) {
		// Reset token is valid, you can proceed with password change
		if err := dbUser.ComparePassword(req.OldPassword); err != nil {
			return errors.New("wrong password")
		}
	} else {
		return errors.New("reset token is invalid or expired")
	}

	if len(req.NewPassword) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	dbUser.Password = req.NewPassword
	if err := dbUser.HashPassword(); err != nil {
		return errors.New("error hashing password")
	}

	_, err = userService.userCollection.UpdateOne(userService.context, bson.M{"_id": userID}, bson.M{"$set": bson.M{"password": dbUser.Password, "reset_token": "", "reset_token_expiry": nil}})
	if err != nil {
		return err
	}

	return nil
}

func (userService *UserAuthService) ForgotPassword(req *requests.ForgotPasswordRequest) error {
	var user models.User
	err := userService.userCollection.FindOne(context.Background(), bson.M{"email": req.Email}).Decode(&user)
	if err != nil {
		return errors.New("email not found")
	}

	// Generate reset token
	resetToken, err := user.GenerateJWT()
	if err != nil {
		return errors.New("error generating reset token")
	}

	// Set reset token and expiry in the database
	update := bson.M{
		"$set": bson.M{
			"reset_token":        resetToken,
			"reset_token_expiry": time.Now().Add(1 * time.Hour),
		},
	}
	_, err = userService.userCollection.UpdateOne(context.Background(), bson.M{"_id": user.Id}, update)
	if err != nil {
		return errors.New("error setting reset token")
	}

	// Send reset email
	mailService := NewMailService()
	resetLink := "http://localhost:8080/reset-password?token=" + resetToken
	body := "Click <a href='" + resetLink + "'>here</a> to reset your password"
	err = mailService.SendEmail(user.Email, "Reset Password", body)
	if err != nil {
		return errors.New("error sending reset email")
	}

	return nil
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
