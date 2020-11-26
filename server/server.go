package server

import (
	"encoding/json"
	"fmt"
	"go-users-crud/database"
	"io/ioutil"
	"net/http"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUser func
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		w.WriteHeader(400)
		w.Write([]byte("Fail to read request body"))
		return
	}

	var user user
	if error = json.Unmarshal(requestBody, &user); error != nil {
		w.WriteHeader(400)
		w.Write([]byte("Fail to convert request body to user struct"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		w.WriteHeader(500)
		w.Write([]byte("Fail to connect on database"))
		return
	}
	defer db.Close()

	statement, error := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
	if error != nil {
		w.WriteHeader(500)
		w.Write([]byte("Fail to prepare query statement"))
		return
	}
	defer statement.Close()

	insert, error := statement.Exec(user.Name, user.Email)
	if error != nil {
		w.WriteHeader(500)
		w.Write([]byte("Fail to insert user on database"))
		return
	}

	idInserted, error := insert.LastInsertId()
	if error != nil {
		w.WriteHeader(500)
		w.Write([]byte("Fail to return the last insert ID"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User created with success! ID: %d", idInserted)))
}

// SearchUsers func
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	db, error := database.Connect()
	if error != nil {
		w.WriteHeader(500)
		w.Write([]byte("Fail to connect on database"))
		return
	}
	defer db.Close()

	lines, error := db.Query("SELECT * FROM users")
	if error != nil {
		w.WriteHeader(500)
		w.Write([]byte("Fail to search users on database"))
		return
	}
	defer lines.Close()

	var users []user
	for lines.Next() {
		var user user

		if error := lines.Scan(&user.ID, &user.Name, &user.Email); error != nil {
			w.WriteHeader(500)
			w.Write([]byte("Fail to scan users from database"))
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if error := json.NewEncoder(w).Encode(users); error != nil {
		w.WriteHeader(500)
		w.Write([]byte("Fail to convert users to JSON"))
		return
	}
}

// SearchUser func
func SearchUser(w http.ResponseWriter, r *http.Request) {

}
