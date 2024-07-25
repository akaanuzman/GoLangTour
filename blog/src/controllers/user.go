// controllers/user.go
package controllers

import (
	"blog/src/db"
	"blog/src/models"
	"context"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func init() {
	client := db.ConnectDB()
	userCollection = client.Database("blog").Collection("users")
}

func Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if user.Email == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Email and password are required"})
	}

	if len(user.Password) < 6 {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Password must be at least 6 characters long"})
	}

	// Validate email format
	if err := models.ValidateEmail(user.Email); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	// Check if email is already taken
	var existingUser models.User
	err := userCollection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Email is already in use"})
	}

	// Hash the password
	if err := user.HashPassword(); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error hashing password"})
	}

	user.Id = primitive.NewObjectID()
	_, err = userCollection.InsertOne(context.Background(), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}

func Login(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	var dbUser models.User
	err := userCollection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&dbUser)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Wrong password or email"})
	}

	if err := dbUser.ComparePassword(user.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Wrong password or email"})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
		"user":  dbUser,
	})
}
