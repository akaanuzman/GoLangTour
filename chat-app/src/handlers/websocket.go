package handlers

import (
	"chat-app/src/core/config"
	"chat-app/src/core/db"
	"chat-app/src/models"
	"chat-app/src/utils"
	"context"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var hashManager = utils.NewHashManager()

// WebSocketHandler handles WebSocket connections and broadcasts
type WebSocketHandler struct {
	upgrader          websocket.Upgrader
	clients           map[*websocket.Conn]bool
	broadcast         chan models.Message
	mutex             *sync.Mutex
	messageCollection *mongo.Collection
	userCollection    *mongo.Collection
}

// NewWebSocketHandler creates a new WebSocketHandler
func NewWebSocketHandler(cfg *config.Config) *WebSocketHandler {
	db := db.Database{}
	db.ConnectDB(cfg)
	messageCollection := db.GetCollection("messages")
	userCollection := db.GetCollection("users")

	return &WebSocketHandler{
		upgrader:          websocket.Upgrader{},
		clients:           make(map[*websocket.Conn]bool),
		broadcast:         make(chan models.Message),
		mutex:             &sync.Mutex{},
		messageCollection: messageCollection,
		userCollection:    userCollection,
	}
}

// HandleConnection handles WebSocket connections
func (wsh *WebSocketHandler) HandleConnections(c echo.Context) error {
	ws, err := wsh.upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	wsh.mutex.Lock()
	wsh.clients[ws] = true
	wsh.mutex.Unlock()

	for {
		var msg models.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			wsh.mutex.Lock()
			delete(wsh.clients, ws)
			wsh.mutex.Unlock()
			break
		}

		// Save the message to MongoDB
		msg.Content = hashManager.HashMessage(msg.Content)
		msg.ID = primitive.NewObjectID()
		msg.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
		_, err = wsh.messageCollection.InsertOne(context.TODO(), msg)
		if err != nil {
			return err
		}

		// Populate sender and receiver information
		sender, err := wsh.getUserByID(msg.SenderID)
		if err != nil {
			return err
		}
		receiver, err := wsh.getUserByID(msg.ReceiverID)
		if err != nil {
			return err
		}

		msg.Sender = sender
		msg.Receiver = receiver

		wsh.broadcast <- msg
	}
	return nil
}

// HandleMessages broadcasts incoming messages to all clients
func (wsh *WebSocketHandler) HandleMessages() {
	for {
		msg := <-wsh.broadcast

		wsh.mutex.Lock()
		for client := range wsh.clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(wsh.clients, client)
			}
		}
		wsh.mutex.Unlock()
	}
}

// getUserByID fetches a user by their ID from the MongoDB users collection
func (wsh *WebSocketHandler) getUserByID(userID primitive.ObjectID) (*models.User, error) {
	var user models.User
	filter := bson.M{"_id": userID}
	err := wsh.userCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
