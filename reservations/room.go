// Python >>>>>>>>> Go

package reservations

import (
	"LESGO/db"
	"fmt"
	"os"
	"strings"
	"time"
)

type Room struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}

const ERROR = "Erreur lors de la récupération des réservations"

type Reservation struct {
	Id       int    `json:"id"`
	RoomId   int    `json:"room_id"`
	Date     string `json:"date"`
	RoomName string `json:"room_name"`
}

func ConvertStringToDatetime(value string) time.Time {
	layout := "2006-01-02 15:04"
	date, err := time.Parse(layout, strings.Replace(value, "T", " ", 1))
	if err != nil {
		fmt.Println(err)
		panic("TEMPORAIRE" + value) // check si la string est bien faite
	}
	return date
}

func ConvertDatetimeToString(datetime time.Time) string {
	return datetime.Format("2006-01-02 15:04")
}

func ListRooms() []Room {
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

func CreateReservation(id int, date string) int {
	db, _ := db.Connect(os.Getenv("USER"), os.Getenv("PASSWORD"))
	defer db.Close()
	newDate := ConvertStringToDatetime(date)
	_, err := db.Exec("INSERT INTO reservation (room_id, date) VALUES (?, ?)", id, newDate)
	if err == nil {
		return 1
	}
	return 0
}

func DeleteReservation(id int) int {
	db, _ := db.Connect(os.Getenv("USER"), os.Getenv("PASSWORD"))
	defer db.Close()
	_, err := db.Exec("DELETE FROM reservation WHERE id =", id)
	if err == nil {
		return 1
	}
	return 0
}

func ListReservations() []Reservation {
	db, _ := db.Connect(os.Getenv("USER"), os.Getenv("PASSWORD"))
	defer db.Close()
	var reservations []Reservation
	rows, err := db.Query("SELECT id, room_id, date FROM reservation")
	if err != nil {
		fmt.Println(ERROR)
		fmt.Println(err)
		return reservations
	}
	var reserv Reservation
	for rows.Next() {
		err = rows.Scan(&reserv.Id, &reserv.RoomId, &reserv.Date)
		if err != nil {
			fmt.Println(ERROR)
			fmt.Println(err)
			return reservations
		}
		rows2, err := db.Query("SELECT name FROM room WHERE id = ?", reserv.RoomId)
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
	db, _ := db.Connect(os.Getenv("USER"), os.Getenv("PASSWORD"))
	defer db.Close()
	var Reservations []Reservation
	val := ConvertStringToDatetime(date)
	rows, _ := db.Query("SELECT id, room_id, date FROM reservation WHERE date = ?", val)
	for rows.Next() {
		var reservation Reservation
		err := rows.Scan(&reservation.Id, &reservation.RoomId, &reservation.Date)
		if err != nil {
			fmt.Println(ERROR)
			fmt.Println(err)
			return Reservations
		}
		rows2, err := db.Query("SELECT name FROM room WHERE id = ?", reservation.RoomId)
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
	db, _ := db.Connect(os.Getenv("USER"), os.Getenv("PASSWORD"))
	defer db.Close()
	var Reservations []Reservation
	rows, _ := db.Query("SELECT id, room_id, date FROM reservation WHERE room_id = ?", id)
	for rows.Next() {
		var reservation Reservation
		err := rows.Scan(&reservation.Id, &reservation.RoomId, &reservation.Date)
		if err != nil {
			fmt.Println(ERROR)
			fmt.Println(err)
			return Reservations
		}
		rows2, err := db.Query("SELECT name FROM room WHERE id = ?", id)
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
