package server

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"ascii-art-web/functions"
)

/*
lets  parse our templates and store them in a struct for
multiple use and avoiding template parrsing caching
*/

type ParsedFiles struct {
	AllTemplates *template.Template
	buf          bytes.Buffer
	lastAsciiArt string
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
	parsedFiles.AllTemplates, err = template.ParseGlob("./templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
}

/*
this the home function the  main entry point for our server
that  will handle the http request only with get request and return the response
*/

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		parsedFiles.AllTemplates.ExecuteTemplate(w, "error.html", "MethodNotAllowed")
		return
	}

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		parsedFiles.AllTemplates.ExecuteTemplate(w, "error.html", "NOT FOUND")
		return
	}
	parsedFiles.buf.Reset()
	err := parsedFiles.AllTemplates.ExecuteTemplate(&parsedFiles.buf, "index.html", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		parsedFiles.AllTemplates.ExecuteTemplate(w, "error.html", "Internal Server Error")
		return
	}
	parsedFiles.AllTemplates.ExecuteTemplate(w, "index.html", nil)
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
		parsedFiles.AllTemplates.ExecuteTemplate(w, "error.html", "Method Not Allowed")
		return
	}
	// extract the input text and ther banner fromthe request
	inputText := r.FormValue("inputText")
	banner := r.FormValue("choice")

	if len(inputText) > maxInputTextLength {
		w.WriteHeader(http.StatusBadRequest)
		parsedFiles.AllTemplates.ExecuteTemplate(w, "error.html", "input text exceeds 500 characters")
		return
	}

	if len(inputText) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		parsedFiles.AllTemplates.ExecuteTemplate(w, "error.html", "Enter a text")
		return
	}

	// Extract out templtae
	templ, err, status := ReadBannerTemplate(banner)
	if err != nil {
		if status {
			w.WriteHeader(http.StatusInternalServerError)
			parsedFiles.AllTemplates.ExecuteTemplate(w, "error.html", "Internal Server Error")
			return
		} else {
			w.WriteHeader(http.StatusNotFound)
			parsedFiles.AllTemplates.ExecuteTemplate(w, "error.html", " BANNER NOT FOUND")
			return
		}
	}

	// This condition for internal errors if the banners get changed
	if len(templ) != 856 {
		w.WriteHeader(http.StatusInternalServerError)
		parsedFiles.AllTemplates.ExecuteTemplate(w, "error.html", "Internal Server Error")
		return
	}

	// Generate our ascii art
	treatedText, err := functions.TraitmentData(templ, inputText)
	parsedFiles.lastAsciiArt = treatedText
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		parsedFiles.AllTemplates.ExecuteTemplate(w, "error.html", "Our ascii program  do not suport non-ascii printable characters")
		return
	}

	// Exexute  the template with the generated ascii art
	// treatedText = strings.TrimSpace(treatedText)
	parsedFiles.buf.Reset()
	err = parsedFiles.AllTemplates.ExecuteTemplate(&parsedFiles.buf, "index.html", treatedText)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		parsedFiles.AllTemplates.ExecuteTemplate(w, "error.html", "Internal Server Error")
		return
	}

	parsedFiles.AllTemplates.ExecuteTemplate(w, "index.html", treatedText)
}

func ServStatic(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		parsedFiles.AllTemplates.ExecuteTemplate(w, "error.html", "Method not allowed")
		return
	}
	file, err := os.Stat(r.URL.Path[1:])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		parsedFiles.AllTemplates.ExecuteTemplate(w, "error.html", "NOT FOUND")
		return
	}
	if file.IsDir() {
		w.WriteHeader(http.StatusNotFound)
		parsedFiles.AllTemplates.ExecuteTemplate(w, "error.html", "NOT FOUND")
		return
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}

func ExportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		parsedFiles.AllTemplates.ExecuteTemplate(w, "error.html", "Method Not Allowed")
		return
	}
	export_result := r.URL.Query().Get("exportme")

	w.Header().Set("Content-Disposition", "attachment; filename=Ascii-Art.txt")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.Itoa(len([]byte(export_result))))

	_, err := w.Write([]byte(export_result))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		parsedFiles.AllTemplates.ExecuteTemplate(w, "error.html", "Internal Server Error")
		return
	}
}
