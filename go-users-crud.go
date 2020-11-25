package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConnection := "golang:golang@/users?charset=utf8&parseTime=True&loc=Local"

	db, error := sql.Open("mysql", stringConnection)

	if error != nil {
		log.Fatal(error)
	}

	defer db.Close()

	if error = db.Ping(); error != nil {
		log.Fatal(error)
	}

	fmt.Println("Database Connection Open")

	users, error := db.Query("select * from users")

	if error != nil {
		log.Fatal(error)
	}

	defer users.Close()

	fmt.Println(users)
}
