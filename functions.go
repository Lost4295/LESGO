package main

import (
	"LESGO/db"
	res "LESGO/reservations"
	"LESGO/web"
	"bufio"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"slices"
	"strconv"
	"strings"
)

const INFO = "\n%sPour quitter, entrez '0' %s"

func HandleErr(err error) {
	if err != nil {
		fmt.Println("Erreur : ", err)
	}
}
func Clear() {
	fmt.Print("\033[H\033[2J")
}

func Start() {
	var verbose, clear, noWeb, noCli bool = false, false, false, false
	if slices.Contains(os.Args, "-v") || slices.Contains(os.Args, "--verbose") {
		verbose = true
	}
	if slices.Contains(os.Args, "-c") || slices.Contains(os.Args, "--clear") {
		clear = true
	}
	if slices.Contains(os.Args, "-nw") || slices.Contains(os.Args, "--no-web") {
		noWeb = true
	}
	if slices.Contains(os.Args, "-nc") || slices.Contains(os.Args, "--no-cli") {
		noCli = true
	}
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	val := strconv.FormatBool(verbose)
	check := os.Getenv("VERBOSE")
	if check != val && verbose {
		fmt.Println("Setting VERBOSE to ", val)
		os.Setenv("VERBOSE", val)
	} else {
		verbose, _ = strconv.ParseBool(check)
	}

	val = strconv.FormatBool(noCli)
	check = os.Getenv("NO_CLI")
	if check != val && noCli {
		fmt.Println("Setting NO_CLI to ", val)
		os.Setenv("NO_CLI", val)
	} else {
		noCli, _ = strconv.ParseBool(check)
	}

	val = strconv.FormatBool(clear)
	check = os.Getenv("CLEAR")
	if check != val && clear{
		fmt.Println("Setting CLEAR to ", val)
		os.Setenv("CLEAR", val)
	} else {
		clear, _ = strconv.ParseBool(check)
	}

	val = strconv.FormatBool(noWeb)
	check = os.Getenv("NO_WEB")
	if check != val && noWeb {
		fmt.Println("Setting NO_WEB to ", val)
		os.Setenv("NO_WEB", val)
	} else {
		noWeb, _ = strconv.ParseBool(check)
	}
	if !noWeb {
		web.Main()
	}
	db.CheckConnection()
}

func ShowMenu() {
	fmt.Println()
	fmt.Println("1. Lister toutes les salles")
	fmt.Println("2. Lister les salles disponibles")
	fmt.Println("3. Créer une réservation")
	fmt.Println("4. Annuler une réservation")
	fmt.Println("5. Visualiser les réservations")
	fmt.Println("6. Visualiser les réservations d'une salle")
	fmt.Println("7. Exporter les réservations")
	fmt.Println("8. Importer des réservations")
	fmt.Println("0. Quitter")
	fmt.Println()
}

func askForYear(scanner *bufio.Scanner) (string, error) {
	var year string
	good := false
	for !good {
		fmt.Printf("\n%sEntrez l'année de réservation (yyyy):%s", GREEN, END)
		fmt.Printf(INFO, BLUE, END)
		scanner.Scan()
		year = scanner.Text()
		if year == "0" {
			return "0", errors.New("Annulation")
		}
		if len(year) == 4 {
			if year[0] == '2' && year[1] == '0' {
				good = true
			} else {
				fmt.Printf("\n%sL'année doit commencer par 20%s", RED, END)
			}
		} else {
			fmt.Printf("\n%sL'année doit être au format yyyy%s", RED, END)
		}
	}
	return year, nil
}
func askForMonth(scanner *bufio.Scanner) (string, error) {
	var month string
	good := false
	for !good {
		fmt.Printf("\n%sEntrez le mois de réservation (mm):%s", GREEN, END)
		fmt.Printf(INFO, BLUE, END)
		scanner.Scan()
		month = scanner.Text()
		if month == "0" {
			return "0", errors.New("Annulation")
		}
		intMonth, err := strconv.Atoi(month)
		HandleErr(err)
		if 01 <= intMonth && intMonth <= 12 {
			good = true
			if len(month) == 1 {
				month = "0" + month
			}
		} else {
			fmt.Printf("\n%sLe mois doit être compris entre 1 et 12%s", RED, END)
		}
	}
	return month, nil

}
func askForDay(month string, year string, scanner *bufio.Scanner) (string, error) {
	var day string
	good := false
	for !good {
		fmt.Printf("\n%sEntrez le jour de réservation (dd):%s", GREEN, END)
		fmt.Printf(INFO, BLUE, END)
		scanner.Scan()
		day = scanner.Text()
		if day == "0" {
			return "0", errors.New("Annulation")
		}
		intDay, err := strconv.Atoi(day)
		HandleErr(err)
		intMonth, err := strconv.Atoi(month)
		HandleErr(err)
		intYear, err := strconv.Atoi(year)
		HandleErr(err)
		switch intMonth {
		case 1, 3, 5, 7, 8, 10, 12:
			if intDay >= 1 && intDay <= 31 {
				good = true
			} else {
				fmt.Printf("\n%sLe jour doit être compris entre 1 et 31%s", RED, END)
			}
		case 4, 6, 9, 11:
			if intDay >= 1 && intDay <= 30 {
				good = true
			} else {
				fmt.Printf("\n%sLe jour doit être compris entre 1 et 30%s", RED, END)
			}
		case 2:
			if intYear%4 == 0 && (intYear%100 != 0 || intYear%400 == 0) {
				if intDay >= 1 && intDay <= 29 {
					good = true
				} else {
					fmt.Printf("\n%sLe jour doit être compris entre 1 et 29%s", RED, END)
				}
			} else {
				if intDay >= 1 && intDay <= 28 {
					good = true
				} else {
					fmt.Printf("\n%sLe jour doit être compris entre 1 et 28%s", RED, END)
				}
			}
		}
		if len(day) == 1 {
			day = "0" + day
		}
	}
	return day, nil
}
func askForHour(scanner *bufio.Scanner) (string, error) {
	var hour string
	good := false
	for !good {
		fmt.Printf("\n%sEntrez l'heure de réservation (hh):%s", GREEN, END)
		fmt.Printf(INFO, BLUE, END)
		scanner.Scan()
		hour = scanner.Text()
		if hour == "0" {
			return "0", errors.New("Annulation")
		}
		intHour, err := strconv.Atoi(hour)
		HandleErr(err)
		if intHour >= 0 && intHour <= 23 {
			good = true
			if len(hour) == 1 {
				hour = "0" + hour
			}
		} else {
			fmt.Printf("\n%sL'heure doit être comprise entre 0 et 23%s", RED, END)
		}
	}
	return hour, nil
}
func askForMinute(scanner *bufio.Scanner) (string, error) {
	var minute string
	good := false
	for !good {
		fmt.Printf("\n%sEntrez les minutes de réservation (mm):%s", GREEN, END)
		fmt.Printf(INFO, BLUE, END)
		scanner.Scan()
		minute = scanner.Text()
		if minute == "0" {
			return "0", errors.New("Annulation")
		}
		intMinute, err := strconv.Atoi(minute)
		HandleErr(err)
		if intMinute >= 0 && intMinute <= 59 {
			good = true
			if len(minute) == 1 {
				minute = "0" + minute
			}
		} else {
			fmt.Printf("%sLes minutes doivent être comprises entre 0 et 59%s", RED, END)
		}
	}
	return minute, nil
}
func CreateDate(scanner *bufio.Scanner) (string, string, string, string, string, error) {

	// Année
	year, err := askForYear(scanner)
	if err != nil {
		return "0", "0", "0", "0", "0", errors.New("Annulation")
	}
	//Mois
	month, err := askForMonth(scanner)
	if err != nil {
		return "0", "0", "0", "0", "0", errors.New("Annulation")
	}
	//Jour
	day, err := askForDay(month, year, scanner)
	if err != nil {
		return "0", "0", "0", "0", "0", errors.New("Annulation")
	}

	//Heure
	hour, err := askForHour(scanner)
	if err != nil {
		return "0", "0", "0", "0", "0", errors.New("Annulation")
	}

	//Minute
	minute, err := askForMinute(scanner)
	if err != nil {
		return "0", "0", "0", "0", "0", errors.New("Annulation")
	}
	return year, month, day, hour, minute, nil
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
