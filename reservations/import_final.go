package reservations

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
	"strconv"
	"strings"
)

func Truncate_seconds(date_time string) string {
	parts := strings.Split(date_time, " ")
	return parts[0] + " " + parts[1][:5]
}

func Import_reserv_from_json(filename string) ([]Reservation, error) {
	/*
		Fonction pour importer des réservations à partir d'un fichier JSON.

		Paramètres :
			- filename (string) : nom du fichier JSON à importer.

		Retourne :
			[]Reservation : tableau des réservations importées.
			error : toute erreur rencontrée lors de l'importation, le cas échéant.
	*/

	// Lire les données JSON à partir du fichier
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var reservations []Reservation

	// Décoder les données JSON en tableau de réservations
	if err := json.Unmarshal(data, &reservations); err != nil {
		return nil, err
	}

	// Créer des réservations à partir des données importées
	for _, reservation := range reservations {
		Create_reservation(reservation.Room_id, Convert_string_to_datetime(Truncate_seconds(reservation.Date_debut)), Convert_string_to_datetime(Truncate_seconds(reservation.Date_fin)))
	}

	return reservations, nil
}

func Import_reserv_from_csv(filename string) ([]Reservation, error) {
	/*
		Fonction pour importer des réservations à partir d'un fichier CSV.

		Paramètres :
			- filename (string) : nom du fichier CSV à importer.

		Retourne :
			[]Reservation : tableau des réservations importées.
			error : toute erreur rencontrée lors de l'importation, le cas échéant.
	*/

	// Ouvrir le fichier CSV
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Créer un lecteur CSV
	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	var reservations []Reservation

	// Lire l'en-tête CSV
	columns, err := reader.Read()
	if err != nil {
		return nil, err
	}

	// Créer une carte pour mapper les colonnes
	column_map := make(map[string]int)
	for i, column := range columns {
		column_map[column] = i
	}

	// Lire les lignes CSV et créer des réservations
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		// Extraire les données de la ligne
		id_str := line[column_map["Id"]]
		room_id_str := line[column_map["RoomId"]]
		date_str := line[column_map["DateDebut"]]
		date_str2 := line[column_map["DateFin"]]
		room_name := line[column_map["RoomName"]]

		// Convertir les données
		id, _ := strconv.Atoi(id_str)
		room_id, _ := strconv.Atoi(room_id_str)
		datet := Truncate_seconds(date_str)
		date2t := Truncate_seconds(date_str2)
		date := Convert_string_to_datetime(datet)
		date2 := Convert_string_to_datetime(date2t)

		// Créer une réservation
		Create_reservation(room_id, date, date2)

		reservation := Reservation{
			Id:        id,
			Room_id:    room_id,
			Date_debut: date_str,
			Date_fin:   date_str2,
			Room_name:  room_name,
		}

		reservations = append(reservations, reservation)
	}

	return reservations, nil
}
