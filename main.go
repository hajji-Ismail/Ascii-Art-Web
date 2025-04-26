package main

import (
	"fmt"
	"log"
	"net/http"

	"ascii-art-web/server"
)

func main() {
	http.HandleFunc("/", server.Home)
	http.HandleFunc("/ascii-art", server.SubmitHandler)
	http.HandleFunc("/css/", server.ServStatic)
	http.HandleFunc("/export", server.ExportHandler)
	fmt.Println("Server is running on port 8080", ">>> http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
