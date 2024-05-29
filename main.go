package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yourusername/myapp/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HandleIndex)
	http.HandleFunc("/add", handlers.HandleAdd)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
