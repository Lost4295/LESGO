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
		panic("TEMPORAIRE")
	}
	return date
}

func IsFree(value string) {
	date := convertStringToDatetime(value)
	db, _ := db.Connect("user", "password")

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
