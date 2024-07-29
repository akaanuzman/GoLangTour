package handlers

import (
	"chat-app/src/core/config"
	"chat-app/src/core/db"
	"chat-app/src/models"
	"context"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// WebSocketHandler handles WebSocket connections and broadcasts
type WebSocketHandler struct {
	upgrader          websocket.Upgrader
	clients           map[*websocket.Conn]bool
	broadcast         chan models.Message
	mutex             *sync.Mutex
	messageCollection *mongo.Collection
}

// NewWebSocketHandler creates a new WebSocketHandler
func NewWebSocketHandler(cfg *config.Config) *WebSocketHandler {
	db := db.Database{}
	db.ConnectDB(cfg)
	collection := db.GetCollection("messages")

	return &WebSocketHandler{
		upgrader:          websocket.Upgrader{},
		clients:           make(map[*websocket.Conn]bool),
		broadcast:         make(chan models.Message),
		mutex:             &sync.Mutex{},
		messageCollection: collection,
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
		// Add additional fields to the message
		msg.ID = primitive.NewObjectID()
		msg.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
		_, err = wsh.messageCollection.InsertOne(context.TODO(), msg)
		if err != nil {
			return err
		}

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
