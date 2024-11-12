package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Message represents a chat message
type Message struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Content    string             `bson:"content" json:"content"`
	SenderID   primitive.ObjectID `bson:"senderId" json:"senderId"`
	ReceiverID primitive.ObjectID `bson:"receiverId" json:"receiverId"`
	CreatedAt  primitive.DateTime `bson:"createdAt" json:"createdAt"`
	Sender     *User              `bson:"-" json:"sender"`
	Receiver   *User              `bson:"-" json:"receiver"`
}
