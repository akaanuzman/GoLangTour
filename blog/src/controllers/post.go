package controllers

import (
	"blog/src/db"
	"blog/src/models"
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var postCollection *mongo.Collection

func init() {
	client := db.ConnectDB()
	postCollection = client.Database("blog").Collection("posts")
}

func CreatePost(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	authorID, _ := primitive.ObjectIDFromHex(claims["username"].(string))

	var post models.Post
	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	post.Id = primitive.NewObjectID()
	post.AuthorID = authorID
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	_, err := postCollection.InsertOne(context.Background(), post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, post)
}

func UpdatePost(c echo.Context) error {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var post models.Post
	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	post.UpdatedAt = time.Now()
	update := bson.M{"$set": post}
	_, err := postCollection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, post)
}

func DeletePost(c echo.Context) error {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	_, err := postCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Post deleted"})
}
