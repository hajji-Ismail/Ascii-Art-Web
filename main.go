package main

import (
	"fmt"
	"net/http"

	"ascii-art-web/server"
)

func main() {
	http.HandleFunc("/", server.Home)
	http.HandleFunc("/css/style.css", server.CssHandler)
	http.HandleFunc("/ascii-art", server.SubmitHandler)
	http.HandleFunc("/css/error.css", server.CssErrHundle)

	fmt.Println("Server is running on port 8080", ">>> http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
