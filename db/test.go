package db

import (
	"math/rand"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Connect creates a connection to the database

func CreateTest() {
	db, err := Connect("user", "password")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	db, err = initDB(db)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
	addRoom(db, "Room", rand.Intn(100))
	}	
}

func addRoom(db *sql.DB, name string, capacity int) error {
	_, err := db.Exec("INSERT INTO room (name, capacity) VALUES (?, ?)", name, capacity)
	if err != nil {
		return err
	}
	return nil
}