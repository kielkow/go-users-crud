package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver to conect on MySQL database
)

// Connect with MySQL database
func Connect() (*sql.DB, error) {
	stringConnection := "golang:golang@/users?charset=utf8&parseTime=True&loc=Local"

	db, error := sql.Open("mysql", stringConnection)

	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		return nil, error
	}

	return db, nil
}
