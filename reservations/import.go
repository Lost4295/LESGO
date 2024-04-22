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
	/*
		Function to import reservations from a JSON file.

		Parameters:
			- filename (string): name of the JSON file to import.

		Returns:
			[]Reservation: array of imported reservations.
			error: any error encountered during import, if any.
	*/

	// Read JSON data from file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var reservations []Reservation

	// Decode JSON data into reservations array
	if err := json.Unmarshal(data, &reservations); err != nil {
		return nil, err
	}

	// Create reservations from imported data
	for _, reservation := range reservations {
		CreateReservation(reservation.RoomId, ConvertStringToDatetime(truncateSeconds(reservation.DateDebut)), ConvertStringToDatetime(truncateSeconds(reservation.DateFin)))
	}

	return reservations, nil
}

func ImportReservFromCSV(filename string) ([]Reservation, error) {
	/*
		Function to import reservations from a CSV file.

		Parameters:
			- filename (string): name of the CSV file to import.

		Returns:
			[]Reservation: array of imported reservations.
			error: any error encountered during import, if any.
	*/

	// Open CSV file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create CSV reader
	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	var reservations []Reservation

	// Read CSV header
	columns, err := reader.Read()
	if err != nil {
		return nil, err
	}

	// Create map to map columns
	columnMap := make(map[string]int)
	for i, column := range columns {
		columnMap[column] = i
	}

	// Read CSV lines and create reservations
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		// Extract data from line
		idStr := line[columnMap["Id"]]
		roomIDStr := line[columnMap["RoomId"]]
		dateStr := line[columnMap["DateDebut"]]
		dateStr2 := line[columnMap["DateFin"]]
		roomName := line[columnMap["RoomName"]]

		// Convert data
		id, _ := strconv.Atoi(idStr)
		roomID, _ := strconv.Atoi(roomIDStr)
		Date := truncateSeconds(dateStr)
		Date2 := truncateSeconds(dateStr2)
		date := ConvertStringToDatetime(Date)
		date2 := ConvertStringToDatetime(Date2)

		// Create reservation
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
