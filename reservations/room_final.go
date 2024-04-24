// Python >>>>>>>>> Go

package reservations

import (
	"LESGO/db"
	"fmt"
	"os"
	"strings"
	"time"
)

// Room represents a room with its properties.
type Room struct {
	id       int    `json:"id"`
	name     string `json:"name"`
	capacity int    `json:"capacity"`
}

// Reservation represents a reservation with its properties.
type Reservation struct {
	id        int    `json:"id"`
	room_id   int    `json:"room_id"`
	date_debut string `json:"date_debut"`
	date_fin   string `json:"date_fin"`
	room_name  string `json:"room_name"`
}

// Constants for error messages and color codes.
const (
	ERROR       = "Error occurred while fetching reservations"
	SELECT      = "SELECT name FROM room WHERE id = ?"
	ERRSELECTED = "The room is already booked during the selected dates."
	RED         = "\033[31;01;51m"
	END         = "\033[0m"
)

func convert_string_to_datetime(value string) time.Time {
	layout := "2006-01-02 15:04"
	date, err := time.Parse(layout, strings.Replace(value, "T", " ", 1))
	if err != nil {
		fmt.Println(err)
		panic(value)
	}
	return date
}

func convert_datetime_to_string(datetime time.Time) string {
	return datetime.Format("2006-01-02 15:04")
}

func list_rooms() []Room {
	/*
		Function to retrieve and return all rooms.

		Returns:
			[]Room: array of Room objects representing rooms.
	*/

	// Connection to the database with the identifiers entered in the .env file
	db, _ := db.Connect(os.Getenv("USER"), os.Getenv("PASSWORD"))
	defer db.Close()
	rows, _ := db.Query("SELECT id, name, capacity FROM room")
	var rooms []Room
	for rows.Next() {
		var room Room
		_ = rows.Scan(&room.id, &room.name, &room.capacity)
		rooms = append(rooms, room)
	}
	return rooms
}

func create_reservation(id int, date time.Time, date2 time.Time) int {
	/*
		Function to create a reservation.

		Parameters:
			- id (int): ID of the room for the reservation.
			- date (time.Time): start date of the reservation.
			- date2 (time.Time): end date of the reservation.

		Returns:
			int: 1 if the reservation is successful, 0 otherwise.
	*/

	db, _ := db.Connect(os.Getenv("USER"), os.Getenv("PASSWORD"))
	defer db.Close()

	// Code to check if the reservation is available on this date
	rows, err := db.Query("SELECT date_debut, date_fin FROM reservation where room_id= ?", id)
	if err != nil {
		fmt.Println(ERROR)
		fmt.Println(err)
		return 0
	}
	for rows.Next() {
		var date_debut string
		var date_fin string
		err = rows.Scan(&date_debut, &date_fin)
		if err != nil {
			fmt.Println(ERROR)
			fmt.Println(err)
			return 0
		}
		if date.Before(convert_string_to_datetime(truncate_seconds(date_debut))) && date2.After(convert_string_to_datetime(truncate_seconds(date_debut))) && date2.Before(convert_string_to_datetime(truncate_seconds(date_fin))) {
			// Display of constants to have color + predefined errors
			fmt.Println(RED, ERRSELECTED, END)
			// Return of 0 to indicate that there was an error
			return 0
		}
		if date.Before(convert_string_to_datetime(truncate_seconds(date_fin))) && date.After(convert_string_to_datetime(truncate_seconds(date_debut))) && date2.After(convert_string_to_datetime(truncate_seconds(date_fin))) {
			fmt.Println(RED, ERRSELECTED, END)
			return 0
		}
		if date.After(convert_string_to_datetime(truncate_seconds(date_debut))) && date2.Before(convert_string_to_datetime(truncate_seconds(date_fin))) {
			fmt.Println(RED, ERRSELECTED, END)
			return 0
		}
	}

	// If no error has been found, the reservation is created in the database
	_, err = db.Exec("INSERT INTO reservation (room_id, date_debut, date_fin) VALUES (?, ?, ?)", id, date, date2)
	if err != nil {
		return 0
	}
	return 1
}

func delete_reservation(id int) int {
	/*
		Function to delete a reservation.

		Parameters:
			- id (int): ID of the reservation to delete.

		Returns:
			int: 1 if the deletion is successful, 0 if the reservation doesn't exist, -1 if an error occurs.
	*/
	db, _ := db.Connect(os.Getenv("USER"), os.Getenv("PASSWORD"))
	defer db.Close()
	rows, err := db.Query("Select id FROM reservation WHERE id = ?", id)
	if err != nil {
		return -1
	}
	if !rows.Next() {
		return 0
	}
	_, err = db.Exec("DELETE FROM reservation WHERE id = ?", id)
	if err != nil {
		return -1
	}
	return 1
}

func list_reservations() []Reservation {
	/*
		Function to retrieve and return all reservations.

		Returns:
			[]Reservation: array of Reservation objects representing reservations.
	*/
	db, _ := db.Connect(os.Getenv("USER"), os.Getenv("PASSWORD"))
	defer db.Close()
	var reservations []Reservation
	rows, err := db.Query("SELECT id, room_id, date_debut, date_fin FROM reservation")
	if err != nil {
		fmt.Println(ERROR)
		fmt.Println(err)
		return reservations
	}
	var reserv Reservation
	for rows.Next() {
		// Retrieves reservation data into a reservation object
		err = rows.Scan(&reserv.id, &reserv.room_id, &reserv.date_debut, &reserv.date_fin)
		if err != nil {
			fmt.Println(ERROR)
			fmt.Println(err)
			return reservations
		}
		// Retrieving the name of the room concerned by the reservation
		rows2, err := db.Query(SELECT, reserv.room_id)
		if err != nil {
			fmt.Println(ERROR)
			fmt.Println(err)
			return reservations
		}
		for rows2.Next() {
			err = rows2.Scan(&reserv.room_name)
			if err != nil {
				fmt.Println(ERROR)
				fmt.Println(err)
				return reservations
			}
		}
		reservations = append(reservations, reserv)
	}
	return reservations
}

func list_reservations_by_date(date string) []Reservation {
	/*
		Function to retrieve and return reservations by date.

		Parameters:
			- date (string): Date to filter reservations.

		Returns:
			[]Reservation: Array of Reservation objects filtered by date.
	*/
	db, _ := db.Connect(os.Getenv("USER"), os.Getenv("PASSWORD"))
	defer db.Close()
	var Reservations []Reservation
	val := convert_string_to_datetime(date)
	rows, _ := db.Query("SELECT id, room_id, date_debut, date_fin FROM reservation WHERE date_debut = ? OR date_fin = ?", val, val)
	for rows.Next() {
		var reservation Reservation
		// Retrieves reservation data into a reservation object
		err := rows.Scan(&reservation.id, &reservation.room_id, &reservation.date_debut, &reservation.date_fin)
		if err != nil {
			fmt.Println(ERROR)
			fmt.Println(err)
			return Reservations
		}
		// Retrieving the name of the room concerned by the reservation
		rows2, err := db.Query(SELECT, reservation.room_id)
		if err != nil {
			fmt.Println(ERROR)
			fmt.Println(err)
			return Reservations
		}
		for rows2.Next() {
			err = rows2.Scan(&reservation.room_name)
			if err != nil {
				fmt.Println(ERROR)
				fmt.Println(err)
				return Reservations
			}
		}
		Reservations = append(Reservations, reservation)
	}
	return Reservations
}
func list_reservations_by_room(id int) []Reservation {
	/*
		Function to retrieve and return reservations by room.

		Parameters:
			- id (int): ID of the room to filter reservations.

		Returns:
			[]Reservation: Array of Reservation objects filtered by room.
	*/
	db, _ := db.Connect(os.Getenv("USER"), os.Getenv("PASSWORD"))
	defer db.Close()
	var Reservations []Reservation
	rows, _ := db.Query("SELECT id, room_id, date_debut, date_fin FROM reservation WHERE room_id = ?", id)
	for rows.Next() {
		var reservation Reservation
		err := rows.Scan(&reservation.id, &reservation.room_id, &reservation.date_debut, &reservation.date_fin)
		if err != nil {
			fmt.Println(ERROR)
			fmt.Println(err)
			return Reservations
		}
		rows2, err := db.Query(SELECT, id)
		if err != nil {
			fmt.Println(ERROR)
			fmt.Println(err)
			return Reservations
		}
		for rows2.Next() {
			err = rows2.Scan(&reservation.room_name)
			if err != nil {
				fmt.Println(ERROR)
				fmt.Println(err)
				return Reservations
			}
		}
		Reservations = append(Reservations, reservation)
	}
	return Reservations
}

func check_salle(id int) int {
	/*
		Function to check if a room exists.

		Parameters:
			- id (int): ID of the room to check.

		Returns:
			int: 1 if the room exists, 0 otherwise.
	*/
	db, _ := db.Connect(os.Getenv("USER"), os.Getenv("PASSWORD"))
	defer db.Close()
	rows, _ := db.Query("SELECT id FROM room WHERE id = ?", id)

	if rows.Next() {
		return 1
	}
	return 0
}
