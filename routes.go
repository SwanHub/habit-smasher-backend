package main

import "net/http"

import "encoding/json"

// AllHabits function displays all current active habits
func AllHabits(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("We out here")
}
