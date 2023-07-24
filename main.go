// main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Person - Our struct for all persons
type Person struct {
	Id           int    `json:"id"`
	Forename     string `json:"forename"`
	Lastname     string `json:"lastname"`
	ProfessionId int    `json:"professionId"`
}

var Persons []Person

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllPersons(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllPersons")
	log.Info("Endpoint Hit: returnAllPersons")
	json.NewEncoder(w).Encode(Persons)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/persons", returnAllPersons)
	log.Fatal(http.ListenAndServe(":4880", myRouter))
}

func main() {
	fmt.Println("Starting: kub-train-go-be-a Endpoint")
	log.Info("Starting: kub-train-go-be-a Endpoint")
	Persons = []Person{
		{Id: 1, Forename: "Bruno", Lastname: "Haferkamp", ProfessionId: 3},
		{Id: 2, Forename: "Claudio", Lastname: "Heysterkamp", ProfessionId: 2},
		{Id: 3, Forename: "Pieter", Lastname: "Koopmans", ProfessionId: 1},
	}
	handleRequests()
}
