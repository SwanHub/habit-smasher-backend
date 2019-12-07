package routes

import (
	"encoding/json"
	"net/http"
)

// AllHabits function displays all current active habits
func AllHabits(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("We out here")
}

// CreateSchema function will be run once, at the initialization of this project to setup the db
func CreateSchema(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("get this db bread")
}
