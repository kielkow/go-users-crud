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
