# User CRUD API (Go)

This project implements a simple RESTful API for managing users using Go.
It supports basic CRUD operations:
- Create User
- Read User(s)
- Update User
- Delete User

This project is designed for machine coding interview rounds and demonstrates:
- REST API design
- In-memory data storage
- HTTP request handling
- JSON encoding/decoding
- Clean and readable code structure
- Basic error handling

---

## Tech Stack
- Language: Go (Golang)
- Framework: net/http
- Storage: In-memory (map)
- Data Format: JSON

---

## User Model

```json
{
  "id": 1,
  "name": "John",
  "email": "john@example.com"
}

| Method | Endpoint    | Description    |
| ------ | ----------- | -------------- |
| POST   | /users      | Create a user  |
| GET    | /users      | Get all users  |
| GET    | /users/{id} | Get user by ID |
| PUT    | /users/{id} | Update user    |
| DELETE | /users/{id} | Delete user    |



Initialize users as empty map
Initialize idCounter = 1

FUNCTION createUser(request):
    Parse request body into user object
    Assign id = idCounter
    Store user in map
    Increment idCounter
    Return created user

FUNCTION getAllUsers():
    Return all users from map

FUNCTION getUserByID(id):
    IF id exists in map:
        Return user
    ELSE:
        Return "User not found"

FUNCTION updateUser(id, request):
    IF id exists:
        Parse request body
        Update name and email
        Return updated user
    ELSE:
        Return "User not found"

FUNCTION deleteUser(id):
    IF id exists:
        Delete user from map
        Return success
    ELSE:
        Return "User not found"



user-crud-api/
│
├── main.go
├── go.mod
└── README.md



package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = make(map[int]User)
var idCounter = 1

func main() {
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/users/", userByIDHandler)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createUser(w, r)
	case http.MethodGet:
		getAllUsers(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func userByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getUserByID(w, r, id)
	case http.MethodPut:
		updateUser(w, r, id)
	case http.MethodDelete:
		deleteUser(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	user.ID = idCounter
	idCounter++
	users[user.ID] = user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	var userList []User
	for _, user := range users {
		userList = append(userList, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userList)
}

func getUserByID(w http.ResponseWriter, r *http.Request, id int) {
	user, exists := users[id]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request, id int) {
	_, exists := users[id]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var updatedUser User
	json.NewDecoder(r.Body).Decode(&updatedUser)
	updatedUser.ID = id
	users[id] = updatedUser

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

func deleteUser(w http.ResponseWriter, r *http.Request, id int) {
	_, exists := users[id]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	delete(users, id)
	w.WriteHeader(http.StatusNoContent)
}


## How to Run

go mod init user-crud-api
go run main.go
