// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Person - Our struct for all persons
type Person struct {
	Id         string `json:"Id"`
	Forename   string `json:"forename"`
	Lastname   string `json:"lastname"`
	Occupation string `json:"occupation"`
}

var Persons []Person

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllPersons(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllPersons")
	json.NewEncoder(w).Encode(Persons)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/persons", returnAllPersons)
	log.Fatal(http.ListenAndServe(":4880", myRouter))
}

func main() {
	Persons = []Person{
		{Id: "1", Forename: "Bruno", Lastname: "Haferkamp", Occupation: "Striezelmeister"},
		{Id: "2", Forename: "Claudio", Lastname: "Heysterkamp", Occupation: "Meyster"},
		{Id: "3", Forename: "Pieter", Lastname: "Koopmans", Occupation: "Fahrer"},
	}
	handleRequests()
}
