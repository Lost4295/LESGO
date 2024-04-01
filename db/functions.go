package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect(user, password string) (*sql.DB, error) {
	db, err := sql.Open("mysql", user+":"+password+"@tcp(localhost:3306)/mydatabase")
	if err != nil {
		return db, err
	}
	return db, nil
}

func createTable(db *sql.DB, table string, queryParams string) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + table + " (id INT AUTO_INCREMENT PRIMARY KEY, " + queryParams + ")")
	if err != nil {
		return err
	}
	return nil
}

func initDB(db *sql.DB) (*sql.DB, error) {
	err := createTable(db, "room", "name VARCHAR(255), capacity INT")
	if err != nil {
		return db, err
	}
	err = createTable(db, "reservation", "room_id INT, date DATETIME")
	if err != nil {
		return db, err
	}
	return db, nil
}
