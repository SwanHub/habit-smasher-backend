package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	// used for the below db setup function, not sure why it shouldn't be included in this package
	_ "github.com/lib/pq"
)

// SetDB establishes db connection, used in all routes
func SetDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err)
	}
	return db
}

// AllHabits function displays all current active habits
func AllHabits(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("We out here")
}

// CreateHabitTable will run once, to initialize the habit db relation
func CreateHabitTable(w http.ResponseWriter, r *http.Request) {
	db := SetDB()
	defer db.Close()

	// db.Exec(`CREATE TABLE habits(habit_id SERIAL PRIMARY KEY, title VARCHAR(100), active BOOLEAN)`)

	db.Exec(`INSERT INTO habits (title, active) VALUES ('biting my nails', true)`)
	db.Exec(`INSERT INTO habits (title, active) VALUES ('watching youtube videos', true)`)
	db.Exec(`INSERT INTO habits (title, active) VALUES ('waking up late', true)`)

	json.NewEncoder(w).Encode("you created and seeded the habits table")
}
