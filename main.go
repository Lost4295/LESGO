package main

import (
	"LESGO/db"
	res "LESGO/reservations"
	"LESGO/web"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func handleErr(err error) {
	if err != nil {
		fmt.Println("Erreur : ", err)
	}
}

func showMenu() {
	fmt.Println()
	fmt.Println("1. Lister toutes les salles")
	fmt.Println("2. Lister les salles disponibles")
	fmt.Println("3. Créer une réservation")
	fmt.Println("4. Annuler une réservation")
	fmt.Println("5. Visualiser les réservations")
	fmt.Println("6. Exporter les réservations")
	fmt.Println("7. Importer des réservations")
	fmt.Println("8. Quitter")
	fmt.Println()
}

const (
	WHITEONRED = ""
	END        = "\033[0m"
	ROUGE      = "\033[31;01;51m"
	GREEN      = "\033[32;01m"
	BLANC      = "\033[37;07m"
)

func main() {
	var verbose bool = false
	if len(os.Args) >= 2 && (os.Args[1] == "-v" || os.Args[1] == "--verbose") {
		verbose = true
	}

	os.Setenv("verbose", strconv.FormatBool(verbose))
	web.Main()
	db.ConnectToDatabase()
	// fmt.Println(time.DateTime)
	scanner := bufio.NewScanner(os.Stdin)
	var number int
	var err error

	for {

		fmt.Println("Bienvenue dans le Service de Réservation en Ligne")
		fmt.Println("-------------------------------------------------")
		showMenu()
		fmt.Print("Sélectionnez une option : ")

		scanner.Scan()
		number, err = strconv.Atoi(scanner.Text())
		if err != nil {
			showMenu()
			fmt.Print("Veuillez entrer un nombre valide : ")
			continue
		}
		if number > 8 || number < 1 {
			showMenu()
			fmt.Print("Veuillez entrer un nombre entre 1 et 8 : ")
			continue
		}

		fmt.Println("Option choisie :", number)
		switch number {
		case 1:
			fmt.Println("Liste des salles :")
			res.ListRooms()
		case 2:
			fmt.Println("Entrez la date de la réservation sous le format yyyy-mm-dd hh:min :")
			scanner.Scan()
			fmt.Println("Liste des salles disponibles :")
			res.AreFree(scanner.Text())
		case 3:
			fmt.Println("Créer une réservation")
			fmt.Print("Entrez le numéro de la salle : ")
			scanner.Scan()
			salle, err := strconv.Atoi(scanner.Text())
			handleErr(err)
			fmt.Print("Entrez la date de la réservation : ")
			scanner.Scan()
			date := scanner.Text()
			res.CreateReservation(salle, date)
		case 4:
			fmt.Println("Annuler une réservation")
			fmt.Print("Entrez le numéro de la réservation : ")
			scanner.Scan()
			id, err := strconv.Atoi(scanner.Text())
			handleErr(err)
			res.DeleteReservation(id)
		case 5:
			fmt.Println("Visualiser les réservations")
			res.ListReservations()
		case 6:
			fmt.Println("Exporter les réservations")
			fmt.Print("Entrez le format de l'export (json/csv) : ")
			scanner.Scan()
			input := scanner.Text()
			inputLower := strings.ToLower(input)

			if inputLower == "json" {
				res.ExportReservToJson("reservations")
			} else if inputLower == "csv" {
				res.ExportReservToCSV("reservations")
			} else {
				fmt.Println("Erreur : Format incorrect")
			}
		case 7:
			fmt.Println("Importer des réservations")
			fmt.Print("Entrez le nom du fichier : ")
			scanner.Scan()
			input := scanner.Text()
			parts := strings.Split(input, ".")

			if parts[len(parts)-1] == "json" {
				res.ImportReservFromJson(input)
			} else if parts[len(parts)-1] == "csv" {
				res.ImportReservFromCSV(input)
			} else {
				fmt.Println("Erreur : Format de fichier incorrect")
			}
		case 8:
			fmt.Println("Quitter")
			os.Exit(0)
		}
		fmt.Println("Appuyer sur n'importe quelle touche pour revenir au menu principal")
		scanner.Scan()
	}
}
