package reservations

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
	"strconv"
	"strings"
)

func truncateSeconds(dateTime string) string {
	parts := strings.Split(dateTime, " ")
	return parts[0] + " " + parts[1][:5]
}

func ImportReservFromJson(filename string) ([]Reservation, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var reservations []Reservation

	if err := json.Unmarshal(data, &reservations); err != nil {
		return nil, err
	}

	for _, reservation := range reservations {
		CreateReservation(reservation.RoomId, ConvertStringToDatetime(truncateSeconds(reservation.DateDebut)), ConvertStringToDatetime(truncateSeconds(reservation.DateFin)))
	}

	return reservations, nil
}

func ImportReservFromCSV(filename string) ([]Reservation, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	var reservations []Reservation

	columns, err := reader.Read()
	if err != nil {
		return nil, err
	}

	columnMap := make(map[string]int)
	for i, column := range columns {
		columnMap[column] = i
	}

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		idStr := line[columnMap["Id"]]
		roomIDStr := line[columnMap["RoomId"]]
		dateStr := line[columnMap["DateDebut"]]
		dateStr2 := line[columnMap["DateFin"]]
		roomName := line[columnMap["RoomName"]]

		id, _ := strconv.Atoi(idStr)
		roomID, _ := strconv.Atoi(roomIDStr)

		Date := truncateSeconds(dateStr)
		Date2 := truncateSeconds(dateStr2)

		date := ConvertStringToDatetime(Date)
		date2 := ConvertStringToDatetime(Date2)
		CreateReservation(roomID, date, date2)

		reservation := Reservation{
			Id:        id,
			RoomId:    roomID,
			DateDebut: Date,
			DateFin:   Date2,
			RoomName:  roomName,
		}

		reservations = append(reservations, reservation)
	}

	return reservations, nil
}
