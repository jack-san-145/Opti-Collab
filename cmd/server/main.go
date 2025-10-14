package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"opti-collab/internal/handlers"
	"os"
)

func main() {
	// Load environment variables
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("No .env file found")
	}

	port := os.Getenv("PORT")

	http.Handle("/", http.FileServer(http.Dir("../../ui"))) // serve HTML
	http.HandleFunc("/ws", handlers.Ws_handler)             // WebSocket endpoint

	log.Println("OptiCollab server running on port", port)
	err = http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		fmt.Println("error while start the go server - ", err)
	}

}
