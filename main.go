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

	fmt.Println("Server on port 8080", ">>> http://localhost:8080")
	http.ListenAndServe("", nil)
}
