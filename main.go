package main

import (
	"LESGO/db"
	res "LESGO/reservations"
	"bufio"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
	"strings"
)

const (
	WHITEONRED = "\033[37;41m"
	END        = "\033[0m"
	RED        = "\033[31;01;51m"
	GREEN      = "\033[32;01m"
	BLANC      = "\033[37;07m"
	BLUE       = "\033[34;01m"
	MAGENTA    = "\033[35;01m"
	CONVERR    = "Erreur de conversion:"
	WHITE      = "\033[37;07m"
)

func main() {
	Start()

	// Si vous réussissez à faire ça, allez y
	// noCli, err := strconv.ParseBool(os.Getenv("NO_CLI"))
	// HandleErr(err)


	os.Exit(0)
	db.CheckConnection()
	scanner := bufio.NewScanner(os.Stdin)
	var number int = -1
	// if !noCli {
	fmt.Printf("%sBienvenue dans le Service de Réservation en Ligne%s", BLUE, END)
	for number != 0 {
		fmt.Printf("\n%s~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~%s", WHITE, END)
		ShowMenu()
		fmt.Printf("%sSélectionnez une option : %s", GREEN, END)
		scanner.Scan()
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			ShowMenu()
			number = -1
			fmt.Printf("%sVeuillez entrer un nombre valide : %s", RED, END)
			continue
		}
		if 0 > number || number > 8 {
			ShowMenu()
			fmt.Printf("%sVeuillez entrer un nombre entre 0 et 8 : %s", RED, END)
			continue
		}

		fmt.Println("Option choisie :", number)
		switch number {
		case 1:
			fmt.Printf("\n%sListe des salles :%s\n", BLUE, END)
			res.ListRooms()
		case 2:
			year, month, day, hour, minute, err := CreateDate(scanner)
			if err != nil {
				fmt.Println("Annulation")
				continue
			}
			fmt.Println("Liste des salles disponibles :")
			res.AreFree(year + "-" + month + "-" + day + " " + hour + ":" + minute)
		case 3:
			fmt.Printf("\n%sCréer une réservation%s\n", BLUE, END)
			fmt.Printf("%sEntrez le numéro de la salle : %s", GREEN, END)
			scanner.Scan()
			salle, err := strconv.Atoi(scanner.Text())
			HandleErr(err)
			if salle == 0 {
				fmt.Println("Annulation")
				continue
			}
			fmt.Print("Entrez la date de la réservation : ")
			if err != nil {
				fmt.Println("Annulation")
				continue
			}
			year, month, day, hour, minute, err := CreateDate(scanner)
			if err != nil {
				fmt.Println("Annulation")
				continue
			}
			res.CreateReservation(salle, year+"-"+month+"-"+day+" "+hour+":"+minute)
		case 4:
			fmt.Println("Annuler une réservation")
			fmt.Print("Entrez le numéro de la réservation : ")
			scanner.Scan()
			id, err := strconv.Atoi(scanner.Text())
			HandleErr(err)
			if id == 0 {
				fmt.Println("Annulation")
				continue
			}
			res.DeleteReservation(id)
		case 5:
			fmt.Println("Visualiser les réservations")
			res.ListReservations()
		case 7:
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
		case 8:
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
		}
	}
	fmt.Println("Quitter")
	fmt.Println("------------------------------------------------------")
	os.Exit(0)
}

// }
