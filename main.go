package main

import (
	"fmt"
	"go-users-crud/server"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", server.CreateUser).Methods(http.MethodPost)

	fmt.Println("Listening port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
