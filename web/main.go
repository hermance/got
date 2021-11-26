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

func addNewHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/addNew" {
		log.Println("404")
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Println("ParseForm() err: %v", err)
		return
	}
	username := r.FormValue("username")
	date := r.FormValue("date")
	log.Println("username = %s\n", username)
	log.Println("date = %s\n", date)
	// TODO api call

	http.Redirect(w, r, "/", http.StatusSeeOther)
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
	mux.HandleFunc("/addNew", addNewHandler)

	http.ListenAndServe(":"+port, mux)
}