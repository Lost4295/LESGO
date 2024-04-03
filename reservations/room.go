// Python >>>>>>>>> Go

package reservations

import (
	"LESGO/db"
	"fmt"
	"time"
)

type Room struct {
	Id       int    `json:id`
	Name     string `json:name`
	Capacity int    `json:capacity`
}

type Reservation struct {
	Id     int       `json:id`
	RoomId int       `json:room_id`
	Date   time.Time `json:date`
}

func convertStringToDatetime(value string) time.Time {
	layout := "2004-09-01 23:04"
	date, err := time.Parse(layout, value)
	if err != nil {
		panic("TEMPORAIRE") // check si la string est bien faite
	}
	return date
}

func IsFree(value string) {
	date := convertStringToDatetime(value)
	db, _ := db.Connect("user", "password")
	defer db.Close()
	rows, _ := db.Query("SELECT roomId FROM Reservation WHERE date !=", date)

	for rows.Next() {
		var roomId int
		_ = rows.Scan(&roomId)
		rows, _ := db.Query("SELECT name, capacity FROM Room WHERE id =", roomId)
		rows.Next()
		var name string
		var capacity int
		_ = rows.Scan(&name, &capacity)
		fmt.Printf("%-15s (Capacité : %d)", name, capacity)
	}
}

func ListRooms() {
	db, _ := db.Connect("user", "password")
	defer db.Close()
	rows, _ := db.Query("SELECT id, name, capacity FROM Room")
	for rows.Next() {
		var id int
		var name string
		var capacity int
		_ = rows.Scan(&id, &name, &capacity)
		fmt.Printf("%d : %-15s (Capacité : %d)", id, name, capacity)
	}
}
func CreateReservation(id int, date string) {
	db, _ := db.Connect("user", "password")
	defer db.Close()
	newDate := convertStringToDatetime(date)
	_, err := db.Exec("INSERT INTO Reservation (roomId, date) VALUES (?, ?)", id, newDate)
	if err != nil {
		fmt.Println("Erreur lors de la création de la réservation")
		return
	} else {
		fmt.Println("Réservation créée avec succès")
	}
}

func DeleteReservation(id int) {
	db, _ := db.Connect("user", "password")
	defer db.Close()
	ListReservations()
	_, err := db.Exec("DELETE FROM Reservation WHERE id =", id)
	if err != nil {
		fmt.Println("Erreur lors de la suppression de la réservation")
		return
	} else {
		fmt.Println("Réservation supprimée avec succès")
	}
}
func ListReservations() {
	db, _ := db.Connect("user", "password")
	rows, _ := db.Query("SELECT id, roomId, date FROM Reservation")
	if rows.Next() {
		for rows.Next() {
			var id int
			var roomId int
			var date time.Time
			_ = rows.Scan(&id, &roomId, &date)
			fmt.Printf("%d : Salle %d réservée le %s", id, roomId, date)
		}
	} else {
		fmt.Println("Aucune réservation")
	}
	defer db.Close()
}
