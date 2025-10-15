package handlers

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

var Rooms = make(map[string][]*websocket.Conn)
var RoomsLock sync.Mutex

func CreateRoom(roomID string) bool {
	RoomsLock.Lock()
	defer RoomsLock.Unlock()

	if _, exists := rooms[roomID]; !exists {
		rooms[roomID] = []*websocket.Conn{}
		return true
	}
	return false
}

func JoinRoom(roomID string, conn *websocket.Conn) error {
	RoomsLock.Lock()
	defer RoomsLock.Unlock()

	if _, exists := rooms[roomID]; !exists {
		return fmt.Errorf("room %s does not exist", roomID)
	}

	rooms[roomID] = append(rooms[roomID], conn)
	return nil
}

func BroadcastToRoom(roomID string, message []byte) {
	RoomsLock.Lock()
	defer RoomsLock.Unlock()

	for _, conn := range rooms[roomID] {
		conn.WriteMessage(websocket.TextMessage, message)
	}
}
