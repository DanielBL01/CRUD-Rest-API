package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"encoding/json"
	"math/rand"
	"strconv"
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

// Get all the schools
func getSchools(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getSchools")

	// Set content type to expect JSON
	w.Header().Set("Content-Type", "application/json")

	// This will present the information in JSON format
	json.NewEncoder(w).Encode(Schools)
}

// Get a single school by ID
func getSchool(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getSchool")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, school := range Schools {
		if school.ID == params["id"] {
			json.NewEncoder(w).Encode(school)
			return
		}
	}

	fmt.Println("Endpoint does not exist")
}

func createSchool(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: createSchool")
	w.Header().Set("Content-Type", "application/json")

	var school School
	err := json.NewDecoder(r.Body).Decode(&school)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	school.ID = strconv.Itoa(rand.Intn(1000000)) // This is not to be used in production
	Schools = append(Schools, school)
	json.NewEncoder(w).Encode(school)
}

func handleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", homepage)
	r.HandleFunc("/schools", getSchools).Methods("GET")
	r.HandleFunc("/schools/{id}", getSchool).Methods("GET")
	r.HandleFunc("/schools", createSchool).Methods("POST")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}

func main() {

	// Initialize a few schools to get started
	Schools = []School {
		School{ID: "1", Name: "University of Waterloo", Location: "Waterloo, Ontario"},
		School{ID: "2", Name: "Harvard University", Location: "Cambridge, Massachusetts"},
	}

	handleRequest()
}

