package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastNane  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Host + req.URL.Path)
	json.NewEncoder(w).Encode(people)
}

func GetPersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	fmt.Println(req.Host + req.URL.Path)

	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Person{})
}

func CreatePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person

	_ = json.NewDecoder(req.Body).Decode(&person)

	person.ID = params["id"]
	people = append(people, person)

	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()

	people = append(people, Person{ID: "1", FirstName: "Valquiria", LastNane: "Bertossi", Address: &Address{City: "SP", State: "SP"}})
	people = append(people, Person{ID: "2", FirstName: "Norberto", LastNane: "Enomoto", Address: &Address{City: "SP", State: "SP"}})
	people = append(people, Person{ID: "3", FirstName: "Kitri", LastNane: "Enomoto", Address: &Address{City: "SP", State: "SP"}})
	people = append(people, Person{ID: "4", FirstName: "Kibana", LastNane: "Enomoto", Address: &Address{City: "SP", State: "SP"}})
	people = append(people, Person{ID: "5", FirstName: "Sophie", LastNane: "Enomoto", Address: &Address{City: "SP", State: "SP"}})

	router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndPoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3001", router))

}
