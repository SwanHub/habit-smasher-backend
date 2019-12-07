package main

import (
	"log"
	"net/http"
	"os"

	"github.com/SwanHub/habit-smasher/routes"
	"github.com/gorilla/mux"
)

// Step 3: setup database with just habits table
// Step 4: query db for habits
// Step 5: Fetch that data from a React App

func main() {
	r := mux.NewRouter()
	// display all active habits
	r.HandleFunc("/habits", routes.AllHabits)
	r.HandleFunc("/schema", routes.CreateHabitTable)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
