package main

import (
	"log"
	"net/http"
	"os"

	"github.com/SwanHub/habit-smasher/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Step 5: Fetch that data from browser
// Step 6: Fetch and display that data from a React App

func main() {
	r := mux.NewRouter().StrictSlash(true)
	// display all active habits
	r.HandleFunc("/habits", routes.AllHabits).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST"}),
		handlers.AllowedOrigins([]string{"*"}))(r)))
}

// r.HandleFunc("/habitsSchema", routes.CreateHabitTable)
