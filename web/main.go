package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

var mainTemplate = template.Must(template.ParseFiles("views/index.html"))
var editTemplate = template.Must(template.ParseFiles("views/new.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	mainTemplate.Execute(w, nil)
}

type New struct {
	Test string
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	editTemplate.Execute(w, &New{Test: "test"})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// routing
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/new", newHandler)

	http.ListenAndServe(":"+port, mux)
}