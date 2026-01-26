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
