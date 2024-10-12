package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"ascii-art-web/functions"
)

type PageData struct {
	Ascii string
}

const maxInputTextLength = 500

var (
	temple01      *template.Template
	ErrorTemplate *template.Template
)

func init() {
	var err error
	temple01, err = template.ParseFiles("template/index.html")
	if err != nil {
		log.Fatal(err)
	}
	ErrorTemplate, err = template.ParseFiles("template/error.html")
	if err != nil {
		log.Fatal(err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		ErrorTemplate.Execute(w, "MethodNotAllowed")
		return
	}

	if r.URL.Path != "/" {

		w.WriteHeader(http.StatusNotFound)
		ErrorTemplate.Execute(w, "NOT FOUND")
		return
	}
	err := temple01.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ErrorTemplate.Execute(w, "Internal Server Error")
		return
	}
}

func ParseForm(r *http.Request) (string, string, error) {
	inputText := r.FormValue("inputText")
	if len(inputText) > maxInputTextLength {
		return "", "", fmt.Errorf("input text exceeds %d characters", maxInputTextLength)
	}
	banner := r.FormValue("choice")
	return inputText, banner, nil
}

func ReadBannerTemplate(banner string) ([]string, error, bool) {
	switch banner {
	case "standard", "shadow", "thinkertoy":
		return functions.ReadFile("banners/" + banner + ".txt")
	default:
		return nil, fmt.Errorf("error: 400 invalid banner choice: %s", banner), false
	}
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		ErrorTemplate.Execute(w, "Method Not Allowed")
		return
	}
	// extract the input text and ther banner fromthe request
	inputText := r.FormValue("inputText")
	banner := r.FormValue("choice")

	if len(inputText) > maxInputTextLength {
		w.WriteHeader(http.StatusBadRequest)
		ErrorTemplate.Execute(w, "input text exceeds 500 characters")
		return
	}

	if len(inputText) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		ErrorTemplate.Execute(w, "Enter a text")
		return
	}
	templ, err, status := ReadBannerTemplate(banner)
	if err != nil {
		if status {
			w.WriteHeader(http.StatusInternalServerError)
			ErrorTemplate.Execute(w, "Internal Server Error")
			return
		} else {
			w.WriteHeader(http.StatusBadRequest)
			ErrorTemplate.Execute(w, "Bad Request")
			return
		}
	}
	treatedText := functions.TraitmentData(templ, inputText)
	err = temple01.Execute(w, PageData{Ascii: treatedText})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ErrorTemplate.Execute(w, "Internal Server Error")
		return
	}
}

func CssHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		ErrorTemplate.Execute(w, "Method Not Allowed")
		return
	}
	http.ServeFile(w, r, "css/style.css")
}

func CssErrHundle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		ErrorTemplate.Execute(w, "Method Not Allowed")
		return
	}
	http.ServeFile(w, r, "css/error.css")
}
