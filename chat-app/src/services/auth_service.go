package services

import (
	"chat-app/src/core/config"
	"chat-app/src/core/db"
	"chat-app/src/core/helpers"
	"chat-app/src/models"
	"chat-app/src/models/requests"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// IAuthService defines the interface for authentication services.
type IAuthService interface {
	// Register registers a new user.
	// Returns an error if the registration fails.
	Register(user *models.User) error

	// Login logs in a user.
	// Returns a JWT token, the user, and an error if the login fails.
	Login(user *models.User) (string, *models.User, error)

	// ChangePassword changes the password of a user.
	// Returns an error if the password change fails.
	ChangePassword(userID primitive.ObjectID, req *requests.PasswordChangeRequest) error

	// ForgotPassword handles the forgot password process.
	// Returns an error if the process fails.
	ForgotPassword(req *requests.ForgotPasswordRequest) error
}

// UserAuthService is a struct that implements the IAuthService interface.
type UserAuthService struct {
	userCollection *mongo.Collection
	context        context.Context
	config         *config.Config
}

// Error messages used in the UserAuthService.
const (
	ErrEmailInUse           = "email is already in use"
	ErrHashingPassword      = "error hashing password"
	ErrWrongPasswordOrEmail = "wrong password or email"
	ErrUserNotFound         = "user not found"
	ErrInvalidToken         = "reset token is invalid or expired"
	ErrPasswordLength       = "password must be at least 6 characters long"
	ErrEmailNotFound        = "email not found"
	ErrGeneratingToken      = "error generating reset token"
	ErrSettingToken         = "error setting reset token"
	ErrSendingEmail         = "error sending reset email"
)

// NewUserService creates a new instance of UserAuthService with the provided database and configuration.
// Returns an instance of IAuthService.
func NewUserAuthService(cfg *config.Config) IAuthService {
	db := db.Database{}
	db.ConnectDB(cfg)
	collection := db.GetCollection("users")
	return &UserAuthService{
		userCollection: collection,
		context:        context.Background(),
		config:         cfg,
	}
}

// Register registers a new user.
// Parameters:
// - user: the user to be registered.
// Returns an error if the registration fails.
func (userService *UserAuthService) Register(user *models.User) error {
	if ok, err := validateEmailAndPassword(user); ok {
		return err
	}

	var existingUser models.User
	err := userService.userCollection.FindOne(userService.context, bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return errors.New(ErrEmailInUse)
	}

	if err := user.HashPassword(); err != nil {
		return errors.New(ErrHashingPassword)
	}

	user.ID = primitive.NewObjectID()
	_, err = userService.userCollection.InsertOne(userService.context, user)
	if err != nil {
		return err
	}

	return nil
}

// Login logs in a user.
// Parameters:
// - user: the user attempting to log in.
// Returns a JWT token, the user, and an error if the login fails.
func (userService *UserAuthService) Login(user *models.User) (string, *models.User, error) {
	var dbUser models.User
	err := userService.userCollection.FindOne(userService.context, bson.M{"email": user.Email}).Decode(&dbUser)
	if err != nil {
		return "", nil, errors.New(ErrWrongPasswordOrEmail)
	}

	if err := dbUser.ComparePassword(user.Password); err != nil {
		return "", nil, errors.New(ErrWrongPasswordOrEmail)
	}

	token, err := dbUser.GenerateJWT()
	if err != nil {
		return "", nil, errors.New(ErrGeneratingToken)
	}

	return token, &dbUser, nil
}

// ChangePassword changes the password of a user.
// Parameters:
// - userID: the ID of the user whose password is to be changed.
// - req: the password change request containing the old and new passwords.
// Returns an error if the password change fails.
func (userService *UserAuthService) ChangePassword(userID primitive.ObjectID, req *requests.PasswordChangeRequest) error {
	var dbUser models.User
	err := userService.userCollection.FindOne(userService.context, bson.M{"_id": userID}).Decode(&dbUser)
	if err != nil {
		return errors.New(ErrUserNotFound)
	}

	if dbUser.ResetToken == "" || dbUser.ResetTokenExpiry.Before(time.Now()) {
		return errors.New(ErrInvalidToken)
	}

	if err := dbUser.ComparePassword(req.OldPassword); err != nil {
		return errors.New(ErrWrongPasswordOrEmail)
	}

	if len(req.NewPassword) < 6 {
		return errors.New(ErrPasswordLength)
	}

	dbUser.Password = req.NewPassword
	if err := dbUser.HashPassword(); err != nil {
		return errors.New(ErrHashingPassword)
	}

	update := bson.M{
		"$set": bson.M{
			"password":           dbUser.Password,
			"reset_token":        "",
			"reset_token_expiry": nil,
		},
	}
	_, err = userService.userCollection.UpdateOne(userService.context, bson.M{"_id": userID}, update)
	if err != nil {
		return err
	}

	return nil
}

// ForgotPassword handles the forgot password process.
// Parameters:
// - req: the forgot password request containing the user's email.
// Returns an error if the process fails.
func (userService *UserAuthService) ForgotPassword(req *requests.ForgotPasswordRequest) error {
	var user models.User
	err := userService.userCollection.FindOne(userService.context, bson.M{"email": req.Email}).Decode(&user)
	if err != nil {
		return errors.New(ErrEmailNotFound)
	}

	resetToken, err := user.GenerateJWT()
	if err != nil {
		return errors.New(ErrGeneratingToken)
	}

	update := bson.M{
		"$set": bson.M{
			"reset_token":        resetToken,
			"reset_token_expiry": time.Now().Add(1 * time.Hour),
		},
	}
	_, err = userService.userCollection.UpdateOne(userService.context, bson.M{"_id": user.ID}, update)
	if err != nil {
		return errors.New(ErrSettingToken)
	}

	mailHelper := helpers.NewMailHelper(userService.config)
	resetLink := "http://localhost:8080/reset-password?token=" + resetToken
	body := "Click <a href='" + resetLink + "'>here</a> to reset your password"
	err = mailHelper.SendEmail(user.Email, "Reset Password", body)
	if err != nil {
		return errors.New(ErrSendingEmail)
	}

	return nil
}

// validateEmailAndPassword validates the email and password of a user.
// Parameters:
// - user: the user whose email and password are to be validated.
// Returns a boolean indicating if there is an error and the error itself.
func validateEmailAndPassword(user *models.User) (bool, error) {
	if user.Email == "" || user.Password == "" {
		return true, errors.New("email and password are required")
	}

	if len(user.Password) < 6 {
		return true, errors.New(ErrPasswordLength)
	}

	if err := models.ValidateEmail(user.Email); err != nil {
		return true, err
	}
	return false, nil
}
