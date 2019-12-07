package routes

// Habit used to unmarshal data from "/habits" route
type Habit struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Active bool   `json:"active"`
}
