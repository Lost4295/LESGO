package reservations

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func export_reserv_to_json(file_name string) {
	/*
		Fonction pour exporter les données de réservation au format JSON

		Paramètres :
			- file_name (string) : nom du fichier JSON qui sera créé avec les données

		Retour :
			Rien
	*/
	reserv := list_reservations()

	// Convertir les données au format JSON
	json_data, _ := json.Marshal(reserv)

	// Création et écriture du fichier destinataire des données JSON
	err := os.WriteFile(file_name+".json", json_data, 0644)
	if err != nil {
		return
	}
	fmt.Println("Fichier d'export des réservations en JSON créé")
}

func export_reserv_to_csv(file_name string) {
	/*
		Fonction pour exporter les données de réservation au format CSV.

		Paramètres :
			- file_name (string) : Nom du fichier CSV qui sera créé avec les données.

		Retour :
			Rien.
	*/
	reserv := list_reservations()

	// Création du fichier destinataire des données CSV
	file, err := os.Create(file_name + ".csv")
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier CSV :", err)
		return
	}
	defer file.Close()

	// Création d'un objet Writer permettant d'écrire des données au format CSV
	writer := csv.NewWriter(file)
	// Écrit toutes les données mises en mémoire tampon dans le io.Writer sous-jacent
	defer writer.Flush()

	// Écriture de l'en-tête CSV avec les différentes colonnes correspondant aux valeurs de réservation
	header := []string{"Id", "RoomId", "DateDebut", "DateFin", "RoomName"}
	if err := writer.Write(header); err != nil {
		fmt.Println("Erreur lors de l'écriture de l'en-tête CSV :", err)
		return
	}

	// Écriture des réservations une par une dans le fichier CSV
	for _, r := range reserv {
		record := []string{
			strconv.Itoa(r.Id),
			strconv.Itoa(r.RoomId),
			r.DateDebut,
			r.DateFin,
			r.RoomName,
		}
		if err := writer.Write(record); err != nil {
			fmt.Println("Erreur lors de l'écriture des données CSV :", err)
			return
		}
	}

	fmt.Println("Fichier d'export des réservations en CSV créé")
}
