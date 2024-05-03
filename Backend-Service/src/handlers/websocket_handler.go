package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/ABHI2598/Backend-Service/src/models"
)

// WebSocketHandler handles WebSocket connections.
type WebSocketHandler struct {
	connections map[*websocket.Conn]bool
	broadcast   chan models.Job
}

func (wh *WebSocketHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Add the new connection to the map
	wh.connections[conn] = true

	// Send initial jobs data to the new connection
	for _, job := range wh.broadcast {
		conn.WriteJSON(job)
	}

	// Continuously listen for new job updates to broadcast
	for job := range wh.broadcast {
		for conn := range wh.connections {
			conn.WriteJSON(job)
		}
	}
}
