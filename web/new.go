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
	Users []User
}

// A Response struct to map the Entire Response
type User struct {
    Name    string    `json:"name"`
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	getUsers(w)
	var emptyUsers []User
	newTemplate.Execute(w, &New{Test: "test", Users: emptyUsers})
}

func addNewApiCall(username string, date string, status string){
	postBody, _ := json.Marshal(map[string]string{
		"username":  username,
		"date": date,
		"status": status,
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

func getUsers(w http.ResponseWriter){
	
	response, err := http.Get("http://127.0.0.1/users")
    if err != nil {
        log.Println(err.Error())
		return
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Println(err)
		return
    }

	var responseObject []User;
	json.Unmarshal(responseData, &responseObject)
	newTemplate.Execute(w, &New{Test: "test", Users: responseObject})
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
	status := r.FormValue("status")

	addNewApiCall(username, date, status)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}