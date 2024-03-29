package db

import (
	"database/sql"
)

func connect(user, password string) (*sql.DB, error) {
	db, err := sql.Open("mysql", user+":"+password+"@tcp(localhost:3306)/L3C")
	if err != nil {
		return db, err
	}
	return db, nil
	// db.Close
}
