package main

import (
	res "LESGO/reservations"
	"bufio"
	"errors"
	"fmt"
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
	RES        = "%d - Reservation n°%d : Salle %d réservée le %s\n"
)

func main() {
	Start()

	// Si vous réussissez à faire ça, allez y
	// Il faut pouvoir lancer le programme sans le CLI
	// noCli, err := strconv.ParseBool(os.Getenv("NO_CLI"))
	// HandleErr(err)

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
		case 0:
			handleZero()
		case 1:
			handleOne()
		case 2:
			err = handleTwo(scanner)
			if err != nil {
				fmt.Println("Annulation")
				continue
			}
		case 3:
			err = handleThree(scanner)
			if err != nil {
				fmt.Println("Annulation")
				continue
			}
		case 4:
			err = handleFour(scanner)
			if err != nil {
				fmt.Println("Annulation")
				continue
			}
		case 5:
			handleFive()
		case 6:
			err = handleSix(scanner)
			if err != nil {
				fmt.Println("Annulation")
				continue
			}
		case 7:
			handleSeven(scanner)
		case 8:
			handleEight(scanner)
		}
	}
}

func handleZero() {
	fmt.Println("Quitter")
	fmt.Println("------------------------------------------------------")
	os.Exit(0)
}
func handleOne() {
	fmt.Printf("\n%sListe des salles :%s\n", BLUE, END)
	rooms := res.ListRooms()
	for id, room := range rooms {
		fmt.Printf("%d : Salle n°%d- %-15s (Capacité: %d)\n", id, room.Id, room.Name, room.Capacity)
	}
}
func handleTwo(scanner *bufio.Scanner) error {
	year, month, day, hour, minute, err := CreateDate(scanner)
	if err != nil {
		return err
	}
	fmt.Println("Liste des salles disponibles :")
	reservations := res.ListReservationsByDate(year + "-" + month + "-" + day + " " + hour + ":" + minute)
	for id, reservation := range reservations {
		fmt.Printf("%d - Salle %d réservée le %s\n", id, reservation.RoomId, reservation.Date)
	}
	//TODO reformat
	return nil
}
func handleThree(scanner *bufio.Scanner) error {
	fmt.Printf("\n%sCréer une réservation%s\n", BLUE, END)
	fmt.Printf("%sEntrez le numéro de la salle ou 0 pour annuler : %s", GREEN, END)
	scanner.Scan()
	salle, err := strconv.Atoi(scanner.Text())
	HandleErr(err)
	if salle == 0 {
		return errors.New("Annulation")
	}
	fmt.Print("Entrez la date de la réservation : ")
	if err != nil {
		return errors.New("Annulation")
	}
	year, month, day, hour, minute, err := CreateDate(scanner)
	if err != nil {
		return errors.New("Annulation")
	}
	res.CreateReservation(salle, year+"-"+month+"-"+day+" "+hour+":"+minute)
	// TODO Retour de fonction
	return nil
}
func handleFour(scanner *bufio.Scanner) error {
	fmt.Println("Annuler une réservation")
	fmt.Print("Entrez le numéro de la réservation ou 0 pour annuler : ")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	HandleErr(err)
	if id == 0 {
		return errors.New("Annulation")
	}
	res.DeleteReservation(id)
	// TODO Retour de fonction
	return nil
}
func handleFive() {
	fmt.Println("Visualiser les réservations")
	reservations := res.ListReservations()
	for id, reservation := range reservations {
		fmt.Printf(RES, id, reservation.Id, reservation.RoomId, reservation.Date)
	}
	//TODO reformat
}
func handleSix(scanner *bufio.Scanner) error {
	fmt.Println("Visualiser les réservations d'une salle")
	fmt.Println("Voulez-vous filtrer par date ou par salle ?")
	fmt.Println("1. Par date")
	fmt.Println("2. Par salle")
	fmt.Print("Entrez votre choix : ")
	scanner.Scan()
	choix, err := strconv.Atoi(scanner.Text())
	HandleErr(err)
	if choix == 1 {
		year, month, day, hour, minute, err := CreateDate(scanner)
		if err != nil {
			return errors.New("Annulation")
		}
		reservations := res.ListReservationsByDate(year + "-" + month + "-" + day + " " + hour + ":" + minute)
		for id, reservation := range reservations {
			fmt.Printf(RES, id, reservation.Id, reservation.RoomId, reservation.Date)
		} //TODO reformat
	} else if choix == 2 {
		fmt.Print("Entrez le numéro de la salle ou 0 pour annuler : ")
		scanner.Scan()
		id, err := strconv.Atoi(scanner.Text())
		HandleErr(err)
		if id == 0 {
			return errors.New("Annulation")
		}
		reservations := res.ListReservationsByRoom(id)
		for id, reservation := range reservations {
			fmt.Printf(RES, id, reservation.Id, reservation.RoomId, reservation.Date)
		} //TODO reformat
	} else {
		fmt.Println(RED, "Erreur : Choix incorrect", END)
	}
	return nil
}
func handleSeven(scanner *bufio.Scanner) {
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
}
func handleEight(scanner *bufio.Scanner) {
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

// }
