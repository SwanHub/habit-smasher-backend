package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SwanHub/habit-smasher/routes"
	"github.com/gorilla/mux"
)

// Step 3: setup database with just habits table
// Step 4: complete the first thin vertical slice by fetching from a React App

func main() {
	r := mux.NewRouter()
	// display all active habits
	r.HandleFunc("/habits", routes.AllHabits)
	r.HandleFunc("/schema", routes.CreateSeedHabitTable)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}

// SetDB establishes db connection, used in all routes
func SetDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err)
	}
	return db
}
