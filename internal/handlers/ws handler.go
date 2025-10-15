package handlers

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var rooms = make(map[string][]*websocket.Conn)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow requests from all the origins
	},
}

var roomsMutex = sync.Mutex{}

func Ws_handler(w http.ResponseWriter, r *http.Request) {

	roomID := r.URL.Query().Get("room_id")
	fmt.Println("room_id - ", roomID)
	if roomID == "" {
		WriteJSON(w, r, map[string]bool{"created": false})
		return
	} else if _, exists := rooms[roomID]; !exists {
		WriteJSON(w, r, map[string]bool{"created": false})
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}

	roomsMutex.Lock()
	rooms[roomID] = append(rooms[roomID], conn)
	fmt.Printf("New connection in room %s. Total: %d\n", roomID, len(rooms[roomID]))
	roomsMutex.Unlock()

	// Start listening for messages
	go listen_room_msg(roomID, conn)
}

func listen_room_msg(roomID string, conn *websocket.Conn) {
	defer conn.Close()
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error - ", err)
			remove_connection(roomID, conn)
			break
		}

		// Broadcast to all clients in the same room safely
		roomsMutex.Lock()
		conns := rooms[roomID]
		conn_copy := conns
		roomsMutex.Unlock()

		for _, c := range conn_copy {
			if c != conn {
				if err := c.WriteMessage(websocket.TextMessage, msg); err != nil {
					fmt.Println("Write error - ", err)
					remove_connection(roomID, c)
				}
			}
		}

	}
}

func remove_connection(roomID string, conn *websocket.Conn) {
	roomsMutex.Lock()
	defer roomsMutex.Unlock()

	conns := rooms[roomID]
	for i, c := range conns {
		if c == conn {
			rooms[roomID] = append(conns[:i], conns[i+1:]...)
			break
		}
	}

	fmt.Printf("Connection left room %s. Total: %d\n", roomID, len(rooms[roomID]))
}
