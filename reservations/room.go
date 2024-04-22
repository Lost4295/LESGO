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
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}

// Reservation represents a reservation with its properties.
type Reservation struct {
	Id        int    `json:"id"`
	RoomId    int    `json:"room_id"`
	DateDebut string `json:"date_debut"`
	DateFin   string `json:"date_fin"`
	RoomName  string `json:"room_name"`
}

// Constants for error messages and color codes.
const (
	ERROR       = "Error occurred while fetching reservations"
	SELECT      = "SELECT name FROM room WHERE id = ?"
	ERRSELECTED = "The room is already booked during the selected dates."
	RED         = "\033[31;01;51m"
	END         = "\033[0m"
)

func ConvertStringToDatetime(value string) time.Time {
	layout := "2006-01-02 15:04"
	date, err := time.Parse(layout, strings.Replace(value, "T", " ", 1))
	if err != nil {
		fmt.Println(err)
		panic(value)
	}
	return date
}

func ConvertDatetimeToString(datetime time.Time) string {
	return datetime.Format("2006-01-02 15:04")
}

func ListRooms() []Room {
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
		_ = rows.Scan(&room.Id, &room.Name, &room.Capacity)
		rooms = append(rooms, room)
	}
	return rooms
}

func CreateReservation(id int, date time.Time, date2 time.Time) int {
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
		var dateDebut string
		var dateFin string
		err = rows.Scan(&dateDebut, &dateFin)
		if err != nil {
			fmt.Println(ERROR)
			fmt.Println(err)
			return 0
		}
		if date.Before(ConvertStringToDatetime(truncateSeconds(dateDebut))) && date2.After(ConvertStringToDatetime(truncateSeconds(dateDebut))) && date2.Before(ConvertStringToDatetime(truncateSeconds(dateFin))) {
			// Display of constants to have color + predefined errors
			fmt.Println(RED, ERRSELECTED, END)
			// Return of 0 to indicate that there was an error
			return 0
		}
		if date.Before(ConvertStringToDatetime(truncateSeconds(dateFin))) && date.After(ConvertStringToDatetime(truncateSeconds(dateDebut))) && date2.After(ConvertStringToDatetime(truncateSeconds(dateFin))) {
			fmt.Println(RED, ERRSELECTED, END)
			return 0
		}
		if date.After(ConvertStringToDatetime(truncateSeconds(dateDebut))) && date2.Before(ConvertStringToDatetime(truncateSeconds(dateFin))) {
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

func DeleteReservation(id int) int {
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

func ListReservations() []Reservation {
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
		err = rows.Scan(&reserv.Id, &reserv.RoomId, &reserv.DateDebut, &reserv.DateFin)
		if err != nil {
			fmt.Println(ERROR)
			fmt.Println(err)
			return reservations
		}
		rows2, err := db.Query(SELECT, reserv.RoomId)
		if err != nil {
			fmt.Println(ERROR)
			fmt.Println(err)
			return reservations
		}
		for rows2.Next() {
			err = rows2.Scan(&reserv.RoomName)
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

func ListReservationsByDate(date string) []Reservation {
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
	val := ConvertStringToDatetime(date)
	// TODO Voir pour faire un between plutôt sur la journée
	rows, _ := db.Query("SELECT id, room_id, date_debut, date_fin FROM reservation WHERE date_debut = ? OR WHERE date_fin = ?", val)
	for rows.Next() {
		var reservation Reservation
		err := rows.Scan(&reservation.Id, &reservation.RoomId, &reservation.DateDebut)
		if err != nil {
			fmt.Println(ERROR)
			fmt.Println(err)
			return Reservations
		}
		rows2, err := db.Query(SELECT, reservation.RoomId)
		if err != nil {
			fmt.Println(ERROR)
			fmt.Println(err)
			return Reservations
		}
		for rows2.Next() {
			err = rows2.Scan(&reservation.RoomName)
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
func ListReservationsByRoom(id int) []Reservation {
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
	// TODO Voir pour faire un between plutôt sur la journée
	rows, _ := db.Query("SELECT id, room_id, date_debut, date_fin FROM reservation WHERE room_id = ?", id)
	for rows.Next() {
		var reservation Reservation
		err := rows.Scan(&reservation.Id, &reservation.RoomId, &reservation.DateDebut, &reservation.DateFin)
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
			err = rows2.Scan(&reservation.RoomName)
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

func CheckSalle(id int) int {
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
