package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Step 1: create server
// Step 2: connect to Heroku
// Step 3: setup database with just habits table
// Step 4: complete the first thin vertical slice by fetching from a React App

func main() {
	r := mux.NewRouter()
	// display all active habits
	r.HandleFunc("/habits", AllHabits)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
