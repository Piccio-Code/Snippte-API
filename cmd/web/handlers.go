package main

import (
	_ "embed"
	"fmt"
	. "github.com/Piccio-Code/Snippte-API"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	ts, err := template.ParseFS(TemplatesFolder, "*.html", "partials/*.html")

	if err != nil {
		log.Fatal("Parsing has not worked: ", err)
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Fatal("Execution has not worked: ", err)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
