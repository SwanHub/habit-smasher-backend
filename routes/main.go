package routes

import (
	"encoding/json"
	"net/http"

	"github.com/SwanHub/Shipt/Shipt-LocalPractice/main"
)

// AllHabits function displays all current active habits
func AllHabits(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("We out here")
}

// CreateHabitTable will run once, to initialize the habit db relation
func CreateHabitTable(w http.ResponseWriter, r *http.Request) {
	db := main.SetDB()
	defer db.Close()

	db.Exec(`CREATE TABLE habits(habit_id SERIAL PRIMARY KEY, title VARCHAR(100), active BOOLEAN)`)

	db.Exec(`INSERT INTO habits (title, active) VALUES ('biting my nails', true)`)
	db.Exec(`INSERT INTO habits (title, active) VALUES ('watching youtube videos', true)`)
	db.Exec(`INSERT INTO habits (title, active) VALUES ('waking up late', true)`)

	json.NewEncoder(w).Encode("you created and seeded the habits table")
}
