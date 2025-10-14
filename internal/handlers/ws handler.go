package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrade HTTP connection to WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow requests from all the origins
	},
}

// Map of roomID -> list of connections
var rooms = make(map[string][]*websocket.Conn)

func Ws_handler(w http.ResponseWriter, r *http.Request) {

	// Upgrade HTTP request to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	// Read roomID from query param
	roomID := r.URL.Query().Get("room")
	if roomID == "" {
		roomID = "default"
	}

	// Add connection to room
	rooms[roomID] = append(rooms[roomID], conn)
	fmt.Printf("New connection in room %s. Total: %d\n", roomID, len(rooms[roomID]))

	// Listen for messages from this client
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}

		// Broadcast to all clients in the same room
		for _, c := range rooms[roomID] {
			if c != conn {
				err := c.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					fmt.Println("Write error:", err)
				}
			}
		}
	}

	// Remove connection on disconnect
	conns := rooms[roomID]
	for i, c := range conns {
		if c == conn {
			rooms[roomID] = append(conns[:i], conns[i+1:]...)
			break
		}
	}

	fmt.Printf("Connection left room %s. Total: %d\n", roomID, len(rooms[roomID]))
}
