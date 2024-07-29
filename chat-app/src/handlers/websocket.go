package handlers

import (
	"chat-app/src/models"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

// WebSocketHandler handles WebSocket connections and broadcasts
type WebSocketHandler struct {
	upgrader  websocket.Upgrader
	clients   map[*websocket.Conn]bool
	broadcast chan models.Message
	mutex     *sync.Mutex
}

// NewWebSocketHandler creates a new WebSocketHandler
func NewWebSocketHandler() *WebSocketHandler {
	return &WebSocketHandler{
		upgrader:  websocket.Upgrader{},
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan models.Message),
		mutex:     &sync.Mutex{},
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
