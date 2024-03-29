// Python >>>>>>>>> Go

package reservations

import (
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

}

// db, err := sql.Open(driverName: "mysql", dataSourceName: "user: password@tcp(localhost)")
// handleErr (err)

// rows, err := db. Query( query: "SELECT id, name FROM test")
// handleErr(err)

// for rows.Next() {
// 	var id int
// 	var name string
// 	err = rows.Scan(&id, &name)
// 	handleErr(err)
// 	fmt. Println(a: "Ligne :", id, name)
// }

// err = db. Exec( query: "insert into test (name) values (?)", args..: "Margot")
// if err != nil {
// 	log. Fatal(err)
// }
