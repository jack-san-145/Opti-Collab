package main

import (
	"fmt"
	"log"
	"net/http"
	"opti-collab/internal/handlers"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("No .env file found")
	}

	port := os.Getenv("PORT")

	router := chi.NewRouter()

	router.Get("/opti-collab/ws", handlers.Ws_handler)

	router.Get("/opti-collab", handlers.ServeIndex)

	staticPath := "/home/Jack_145/OptiCollab/ui/static" //path which the css,js files lives
	router.Handle("/opti-collab/static*", http.StripPrefix("/opti-collab/static", http.FileServer(http.Dir(staticPath))))

	fmt.Println("server running on " + port + "...")
	err = http.ListenAndServe("0.0.0.0:"+port, router)
	if err != nil {
		fmt.Println("error while start the go server - ", err)
	}

}
