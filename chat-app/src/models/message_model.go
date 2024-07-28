package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Message represents a chat message
type Message struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	SenderID   primitive.ObjectID `json:"senderId" bson:"senderId"`
	ReceiverID primitive.ObjectID `json:"receiverId" bson:"receiverId"`
	Content    string             `json:"content" bson:"content"`
	CreatedAt  primitive.DateTime `json:"createdAt" bson:"createdAt"`
}
