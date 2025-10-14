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

	router.Get("opti-collab/ws", handlers.Ws_handler)

	staticHtmlPath := "/home/Jack_145/OptiCollab/ui/templates" //path which html files live
	router.Handle("/opti-collab/*", http.StripPrefix("/opti-collab/", http.FileServer(http.Dir(staticHtmlPath))))

	staticFilePath := "/home/Jack_145/OptiCollab/ui/static" //path which the css,js files lives
	router.Handle("/opti-collab/static*", http.StripPrefix("/opti-collab/static", http.FileServer(http.Dir(staticFilePath))))

	log.Println("OptiCollab server running on port", port)
	err = http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		fmt.Println("error while start the go server - ", err)
	}

}
