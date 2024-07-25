// models/post.go
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	Content   string             `bson:"content"`
	AuthorID  primitive.ObjectID `bson:"author_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
