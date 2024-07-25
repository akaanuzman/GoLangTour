package services

import (
	"blog/src/db"
	"blog/src/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IPostService interface {
	GetAllPosts() ([]models.Post, error)
	GetPostById(id string) (*models.Post, error)
	CreatePost(post *models.Post) error
	UpdatePost(post *models.Post) error
	DeletePost(id string) error
}

type PostService struct {
	postCollection *mongo.Collection
	context        context.Context
}

func NewPostService() IPostService {
	client := db.ConnectDB()
	return &PostService{
		postCollection: client.Database("blog").Collection("posts"),
		context:        context.Background(),
	}
}

func (postService *PostService) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	cursor, err := postService.postCollection.Find(postService.context, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(postService.context)
	if err := cursor.All(postService.context, &posts); err != nil {
		return nil, err
	}
	return posts, nil
}

func (postService *PostService) GetPostById(id string) (*models.Post, error) {
	var post models.Post
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = postService.postCollection.FindOne(postService.context, bson.M{"_id": objectId}).Decode(&post)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (postService *PostService) CreatePost(post *models.Post) error {
	_, err := postService.postCollection.InsertOne(postService.context, post)
	if err != nil {
		return err
	}
	return nil
}

func (postService *PostService) UpdatePost(post *models.Post) error {
	update := bson.M{"$set": post}
	_, err := postService.postCollection.UpdateOne(postService.context, bson.M{"_id": post.Id}, update)
	if err != nil {
		return err
	}
	return nil
}

func (postService *PostService) DeletePost(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = postService.postCollection.DeleteOne(postService.context, bson.M{"_id": objectId})
	if err != nil {
		return err
	}
	return nil
}
