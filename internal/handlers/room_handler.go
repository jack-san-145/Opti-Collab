package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func GroupCreationHandler(w http.ResponseWriter, r *http.Request) {
	roomID := r.URL.Query().Get("room_id")
	if roomID == "" {
		WriteJSON(w, r, map[string]any{
			"created": false,
			"error":   "room already exists",
		})
		return
	}

	roomsMutex.Lock()
	defer roomsMutex.Unlock()

	if _, exists := rooms[roomID]; exists {
		WriteJSON(w, r, map[string]any{
			"created": false,
			"error":   "room already exists",
		})
		return
	}

	rooms[roomID] = []*websocket.Conn{}
	fmt.Printf("Room %s created\n", roomID)
	WriteJSON(w, r, map[string]bool{"created": true})
}
