package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// User struct representing a user
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// getUsers returns a list of sample users
func getUsers() []User {
	users := []User{
		{ID: 1, Username: "user1", Email: "user1@example.com"},
		{ID: 2, Username: "user2", Email: "user2@example.com"},
		{ID: 3, Username: "user3", Email: "user3@example.com"},
	}
	return users
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	users := getUsers()
	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/users", usersHandler)

	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
