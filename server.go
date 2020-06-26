package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"encoding/json"
)

type School struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Location string `json:"location"`
}

// We need a slice (not a array because we don't want a fixed size)
var Schools []School

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: homepage")
	fmt.Fprintf(w, "Welcome! This is the homepage")
}

func getSchools(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getSchools")

	// This will present the information in JSON format
	json.NewEncoder(w).Encode(Schools)
}

func getSchool(w http.ResponseWriter, r*http.Request) {
	fmt.Println("Endpoint hit: getSchool")
	params := mux.Vars(r)

	for _, school := range Schools {
		if school.ID == params["id"] {
			json.NewEncoder(w).Encode(school)
			return
		}
	}

	fmt.Println("Endpoint does not exist")
}

func handleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", homepage)
	r.HandleFunc("/schools", getSchools).Methods("GET")
	r.HandleFunc("/schools/{id}", getSchool).Methods("GET")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}

func main() {
	Schools = []School {
		School{ID: "1", Name: "University of Waterloo", Location: "Waterloo, Ontario"},
		School{ID: "2", Name: "Harvard University", Location: "Cambridge, Massachusetts"},
	}

	handleRequest()
}

