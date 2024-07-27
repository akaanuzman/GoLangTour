package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Content   string             `bson:"content" json:"content"`
	UserId    primitive.ObjectID `bson:"userId" json:"userId"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
	Tags      []string           `bson:"tags" json:"tags"`
	IsActive  bool               `bson:"isActive" json:"isActive"`
}
