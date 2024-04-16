package reservations

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func ExportReservToJson(fileName string) {
	reserv := ListReservationsReturn()
	jsonData, _ := json.Marshal(reserv)

	err := os.WriteFile(fileName+".json", jsonData, 0644)
	if err != nil {
		return
	}
	fmt.Println("Fichier d'export des réservations en Json créé")
}

func ExportReservToCSV(fileName string) {
	reserv := ListReservationsReturn()
	file, err := os.Create(fileName + ".csv")
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier CSV:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Id", "RoomId", "Date"}
	if err := writer.Write(header); err != nil {
		fmt.Println("Erreur lors de l'écriture de l'en-tête CSV:", err)
		return
	}

	for _, r := range reserv {
		record := []string{
			strconv.Itoa(r.Id),
			strconv.Itoa(r.RoomId),
			// ConvertDatetimeToString(r.Date),
			r.Date,
		}
		if err := writer.Write(record); err != nil {
			fmt.Println("Erreur lors de l'écriture des données CSV:", err)
			return
		}
	}

	fmt.Println("Fichier d'export des réservations en CSV créé")
}
