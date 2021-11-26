package main

import (
	"html/template"
	"log"
	"net/http"
)
var newTemplate = template.Must(template.ParseFiles("views/new.html"))

type New struct {
	Test string
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	newTemplate.Execute(w, &New{Test: "test"})
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