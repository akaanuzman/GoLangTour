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

var userCollection *mongo.Collection = db.ConnectDB().Database("blog").Collection("users")

func Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	user.Id = primitive.NewObjectID()
	_, err := userCollection.InsertOne(context.Background(), user)

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
	err := userCollection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&dbUser)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
