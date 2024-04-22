package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(user string, password string) (*sql.DB, error) {
	/*
		Function to establish a connection to the database.

		Parameters:
			- user (string): Username for database authentication.
			- password (string): Password for database authentication.

		Returns:
			(*sql.DB, error): Database connection object and any error encountered.
	*/

	// Retrieving environment variables for host, port and database name
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	dbname := os.Getenv("DBNAME")
	db, err := sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+dbname)
	if err != nil {
		return db, err
	}
	return db, nil
}

func createTable(db *sql.DB, table string, queryParams string) error {
	/*
		Function to create a table in the database.

		Parameters:
			- db (*sql.DB): Database connection object.
			- table (string): Name of the table to create.
			- queryParams (string): Query parameters for creating the table.

		Returns:
			error: Any error encountered during table creation.
	*/

	// Creation of the table if it does not already exist (avoid overwriting it)
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + table + " (id INT AUTO_INCREMENT PRIMARY KEY, " + queryParams + ")")
	if err != nil {
		return err
	}
	return nil
}

func initDB(db *sql.DB) (*sql.DB, error) {
	/*
		Function to initialize the database with required tables.

		Parameters:
			- db (*sql.DB): Database connection object.

		Returns:
			(*sql.DB, error): Database connection object and any error encountered.
	*/
	err := createTable(db, "room", "name VARCHAR(255), capacity INT")
	if err != nil {
		return db, err
	}
	err = createTable(db, "reservation", "room_id INT, date_debut DATETIME, date_fin DATETIME, FOREIGN KEY (room_id) REFERENCES room(id)")
	if err != nil {
		return db, err
	}
	return db, nil
}

func addRoom(db *sql.DB, name string, capacity int) error {
	/*
		Function to add a room to the database.

		Parameters:
			- db (*sql.DB): Database connection object.
			- name (string): Name of the room.
			- capacity (int): Capacity of the room.

		Returns:
			error: Any error encountered during room addition.
	*/
	_, err := db.Exec("INSERT INTO room (name, capacity) VALUES (?, ?)", name, capacity)
	if err != nil {
		return err
	}
	return nil
}
