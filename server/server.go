package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"ascii-art-web/functions"
)

/*
lets  parse our templates and store them in a struct for
multiple use and avoiding template parrsing caching
*/

type ParsedFiles struct {
	Temple01      *template.Template
	ErrorTemplate *template.Template
	lastAsciiArt  string
}

var parsedFiles ParsedFiles

// this variable for max input text  length that user can  input
const maxInputTextLength = 500

/*
this function init() for initialzing  our parsedFiles struct feilds
to ensure  that our templates are parsed only once
*/
func init() {
	var err error
	parsedFiles.Temple01, err = template.ParseFiles("template/index.html")
	if err != nil {
		log.Fatalf("Error parsing the file: %s", err)
		return
	}
	parsedFiles.ErrorTemplate, err = template.ParseFiles("template/error.html")
	if err != nil {
		log.Fatalf("Error parsing the file: %s", err)
		return
	}
}

/*
this the home function the  main entry point for our server
that  will handle the http request only with get request and return the response
*/

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		parsedFiles.ErrorTemplate.Execute(w, "MethodNotAllowed")
		return
	}

	if r.URL.Path != "/" {

		w.WriteHeader(http.StatusNotFound)
		parsedFiles.ErrorTemplate.Execute(w, "NOT FOUND")
		return
	}

	err := parsedFiles.Temple01.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		parsedFiles.ErrorTemplate.Execute(w, "Internal Server Error")
		return
	}
}

/*
this function for reading the banners depend on the user banner
and return the specific banners as a slice of string and error
if  any error occur
*/
func ReadBannerTemplate(banner string) ([]string, error, bool) {
	switch banner {
	case "standard", "shadow", "thinkertoy":
		return functions.ReadFile("banners/" + banner + ".txt")
	default:
		return nil, fmt.Errorf("error: 400 invalid banner choice: %s", banner), false
	}
}

/*
this function handle  the http request with post method
that post to the client the generated ascii arrt
*/
func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		parsedFiles.ErrorTemplate.Execute(w, "Method Not Allowed")
		return
	}

	// extract the input text and ther banner fromthe request
	inputText := r.FormValue("inputText")
	banner := r.FormValue("choice")

	if len(inputText) > maxInputTextLength {
		w.WriteHeader(http.StatusBadRequest)
		parsedFiles.ErrorTemplate.Execute(w, "input text exceeds 500 characters")
		return
	}

	if len(inputText) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		parsedFiles.ErrorTemplate.Execute(w, "Enter a text")
		return
	}

	// Extract out templtae
	templ, err, status := ReadBannerTemplate(banner)
	if err != nil {
		if status {
			w.WriteHeader(http.StatusInternalServerError)
			parsedFiles.ErrorTemplate.Execute(w, "Internal Server Error")
			return
		} else {
			w.WriteHeader(http.StatusNotFound)
			parsedFiles.ErrorTemplate.Execute(w, " BANNER NOT FOUND")
			return
		}
	}

	// This condition for internal errors if the banners get changed
	if len(templ) != 856 {
		w.WriteHeader(http.StatusInternalServerError)
		parsedFiles.ErrorTemplate.Execute(w, "Internal Server Error")
		return
	}

	// Generate our ascii art
	treatedText, err := functions.TraitmentData(templ, inputText)
	parsedFiles.lastAsciiArt = treatedText
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		parsedFiles.ErrorTemplate.Execute(w, "Our ascii program  do not suport non-ascii printable characters")
		return
	}

	// Exexute  the template with the generated ascii art
	err = parsedFiles.Temple01.Execute(w, treatedText)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		parsedFiles.ErrorTemplate.Execute(w, "Internal Server Error")
		return
	}
}

func ServStatic(w http.ResponseWriter, r *http.Request) {
	// Check if the request is for the CSS directory itself
	if r.URL.Path == "/css/" || r.URL.Path == "/css" {
		w.WriteHeader(http.StatusNotFound)
		parsedFiles.ErrorTemplate.Execute(w, "NOT FOUND")
		return
	}
	_,err := os.Stat(r.URL.Path[1:])
	if  err != nil {
		w.WriteHeader(http.StatusNotFound)
		parsedFiles.ErrorTemplate.Execute(w, "NOT FOUND")
		return
	}

	// Serve the CSS file if it's a valid request
	http.StripPrefix("/css/", http.FileServer(http.Dir("css"))).ServeHTTP(w, r)
}

func ExportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		parsedFiles.ErrorTemplate.Execute(w, "Method Not Allowed")
		return
	}
	export_result := r.URL.Query().Get("exportme")
	
	w.Header().Set("Content-Disposition", "attachment; filename=Ascii-Art.txt")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-Control", "no-cache")
	

	_, err := w.Write([]byte(export_result))
	if err != nil {
		//log.Println("Error writing file:", err)
		w.WriteHeader(http.StatusInternalServerError)
		parsedFiles.ErrorTemplate.Execute(w, "Internal Server Error")
		return
	}
}
