package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
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

func addNewApiCall(username string, date string){
	postBody, _ := json.Marshal(map[string]string{
		"username":  username,
		"date": date,
	 })

	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post("http://127.0.0.1/workday", "application/json", responseBody)
	//Handle Error
	if err != nil {
		// TODO display error to the user
		log.Println("An Error Occured :", err)
		return
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// TODO display error to the user
		log.Println(err)
		return
	}
	sb := string(body)
   	// todo figure what to do with response body
	log.Println(sb)
}

func addNewHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/addNew" {
		log.Println("404")
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Println("ParseForm() err:", err)
		return
	}
	username := r.FormValue("username")
	date := r.FormValue("date")

	addNewApiCall(username, date)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}