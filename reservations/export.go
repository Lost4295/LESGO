package reservations

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func ExportReservToJson(fileName string) {
	/*
		Function to export reservation data in JSON format

		Parameters:
			- fileName (string) : name of the json file that will be created with the data

		Return:
			Nothing
	*/
	reserv := ListReservations()

	// Convert data to json format
	jsonData, _ := json.Marshal(reserv)

	// Creation and writing of the recipient file of the json data
	err := os.WriteFile(fileName+".json", jsonData, 0644)
	if err != nil {
		return
	}
	fmt.Println("Fichier d'export des réservations en Json créé")
}

func ExportReservToCSV(fileName string) {
	/*
		Function to export reservation data in CSV format.

		Parameters:
			- fileName (string): Name of the CSV file that will be created with the data.

		Returns:
			Nothing.
	*/
	reserv := ListReservations()

	// Creation of the recipient file of the csv data
	file, err := os.Create(fileName + ".csv")
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier CSV:", err)
		return
	}
	defer file.Close()

	// Creation of a Writer object allowing you to write data in csv format
	writer := csv.NewWriter(file)
	// Writes all buffered data to the underlying io.Writer
	defer writer.Flush()

	// Writing the csv header with the different columns corresponding to the reservation values
	header := []string{"Id", "RoomId", "DateDebut", "DateFin", "RoomName"}
	if err := writer.Write(header); err != nil {
		fmt.Println("Erreur lors de l'écriture de l'en-tête CSV:", err)
		return
	}

	// Writing reservations one by one in the csv file
	for _, r := range reserv {
		record := []string{
			strconv.Itoa(r.Id),
			strconv.Itoa(r.RoomId),
			r.DateDebut,
			r.DateFin,
			r.RoomName,
		}
		if err := writer.Write(record); err != nil {
			fmt.Println("Erreur lors de l'écriture des données CSV:", err)
			return
		}
	}

	fmt.Println("Fichier d'export des réservations en CSV créé")
}
