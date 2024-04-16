package reservations

import (
	"encoding/csv"
	"encoding/json"
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
		CreateReservationReturn(reservation.RoomId, truncateSeconds(reservation.Date))
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
		if err != nil {
			return nil, err
		}

		idStr := line[columnMap["Id"]]
		roomIDStr := line[columnMap["RoomId"]]
		dateStr := line[columnMap["Date"]]
		roomName := line[columnMap["RoomName"]]

		id, _ := strconv.Atoi(idStr)
		roomID, _ := strconv.Atoi(roomIDStr)

		date := truncateSeconds(dateStr)

		CreateReservationReturn(roomID, date)

		reservation := Reservation{
			Id:       id,
			RoomId:   roomID,
			Date:     date,
			RoomName: roomName,
		}

		reservations = append(reservations, reservation)
	}

	return reservations, nil
}
