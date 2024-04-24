package main

import (
	"LESGO/db"
	res "LESGO/reservations"
	"LESGO/web"
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const INFO = "\n%sPour quitter, entrez '0' %s"
const NORES = "Aucune réservation existante."

func Handle_err(err error) {
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
	if check != val && clear {
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
	db.Check_connection()
}

func Show_menu() {
	fmt.Println()
	fmt.Println("1. Lister toutes les salles")
	fmt.Println("2. Créer une réservation")
	fmt.Println("3. Annuler une réservation")
	fmt.Println("4. Visualiser les réservations")
	fmt.Println("5. Visualiser les réservations d'une salle")
	fmt.Println("6. Exporter les réservations")
	fmt.Println("7. Importer des réservations")
	fmt.Println("0. Quitter")
	fmt.Println()
}

func Ask_for_year(scanner *bufio.Scanner) (string, error) {
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
func Ask_for_month(scanner *bufio.Scanner) (string, error) {
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
		int_month, err := strconv.Atoi(month)
		Handle_err(err)
		if 01 <= int_month && int_month <= 12 {
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
func Ask_for_day(month string, year string, scanner *bufio.Scanner) (string, error) {
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
		int_day, err := strconv.Atoi(day)
		Handle_err(err)
		int_month, err := strconv.Atoi(month)
		Handle_err(err)
		int_year, err := strconv.Atoi(year)
		Handle_err(err)
		switch int_month {
		case 1, 3, 5, 7, 8, 10, 12:
			if int_day >= 1 && int_day <= 31 {
				good = true
			} else {
				fmt.Printf("\n%sLe jour doit être compris entre 1 et 31%s", RED, END)
			}
		case 4, 6, 9, 11:
			if int_day >= 1 && int_day <= 30 {
				good = true
			} else {
				fmt.Printf("\n%sLe jour doit être compris entre 1 et 30%s", RED, END)
			}
		case 2:
			if int_year%4 == 0 && (int_year%100 != 0 || int_year%400 == 0) {
				if int_day >= 1 && int_day <= 29 {
					good = true
				} else {
					fmt.Printf("\n%sLe jour doit être compris entre 1 et 29%s", RED, END)
				}
			} else {
				if int_day >= 1 && int_day <= 28 {
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
func Ask_for_hour(scanner *bufio.Scanner) (string, error) {
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
		Handle_err(err)
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
func Ask_for_minute(scanner *bufio.Scanner) (string, error) {
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
		Handle_err(err)
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
func Create_date(scanner *bufio.Scanner) (string, string, string, string, string, error) {

	// Année
	year, err := Ask_for_year(scanner)
	if err != nil {
		return "0", "0", "0", "0", "0", errors.New("Annulation")
	}
	//Mois
	month, err := Ask_for_month(scanner)
	if err != nil {
		return "0", "0", "0", "0", "0", errors.New("Annulation")
	}
	//Jour
	day, err := Ask_for_day(month, year, scanner)
	if err != nil {
		return "0", "0", "0", "0", "0", errors.New("Annulation")
	}

	//Heure
	hour, err := Ask_for_hour(scanner)
	if err != nil {
		return "0", "0", "0", "0", "0", errors.New("Annulation")
	}

	//Minute
	minute, err := Ask_for_minute(scanner)
	if err != nil {
		return "0", "0", "0", "0", "0", errors.New("Annulation")
	}
	return year, month, day, hour, minute, nil
}
func handle_zero() {
	fmt.Println("Quitter")
	fmt.Println("------------------------------------------------------")
	os.Exit(0)
}
func handle_one() {
	fmt.Printf("\n%sListe des salles :%s\n", BLUE, END)
	rooms := res.List_rooms()
	if len(rooms) == 0 {
		fmt.Println(RED, "Aucune salle existante.", END)
	}
	for id, room := range rooms {
		fmt.Printf("%d : Salle n°%d- %-15s (Capacité: %d)\n", id, room.Id, room.Name, room.Capacity)
	}
}
func handle_two(scanner *bufio.Scanner) error {
	fmt.Printf("\n%sCréer une réservation%s\n", BLUE, END)
	fmt.Printf("%sEntrez le numéro de la salle ou 0 pour annuler : %s", GREEN, END)
	scanner.Scan()
	salle, err := strconv.Atoi(scanner.Text())
	Handle_err(err)
	if salle == 0 {
		return errors.New("Annulation")
	}
	if res.Check_salle(salle) == 0 {
		fmt.Println(RED, "Erreur : Salle inexistante", END)
		return errors.New("Annulation")
	}
	fmt.Print("Entrez la date de début de la réservation : ")
	if err != nil {
		return errors.New("Annulation")
	}
	year, month, day, hour, minute, err := Create_date(scanner)
	if err != nil {
		return errors.New("Annulation")
	}
	date1 := year + "-" + month + "-" + day + " " + hour + ":" + minute
	fmt.Print("Entrez la date de fin de la réservation : ")
	year, month, day, hour, minute, err = Create_date(scanner)
	if err != nil {
		return errors.New("Annulation")
	}
	date2 := year + "-" + month + "-" + day + " " + hour + ":" + minute
	datetime1 := res.Convert_string_to_datetime(date1)
	datetime2 := res.Convert_string_to_datetime(date2)
	if datetime1.After(datetime2) {
		fmt.Println(RED, "Erreur : La date de début doit être avant la date de fin", END)
		return errors.New("Annulation")
	}
	time.Sleep(1 * time.Second)
	if res.Create_reservation(salle, datetime1, datetime2) == 1 {
		fmt.Println(GREEN, "Réservation créée", END)
	} else {
		fmt.Println(RED, "Erreur lors de la création de la réservation", END)
	}
	return nil
}
func handle_three(scanner *bufio.Scanner) error {
	fmt.Println("Annuler une réservation")
	fmt.Print("Entrez le numéro de la réservation ou 0 pour annuler : ")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	Handle_err(err)
	if id == 0 {
		return errors.New("Annulation")
	}
	time.Sleep(1 * time.Second)
	ret := res.Delete_reservation(id)
	if ret == 1 {
		fmt.Println(GREEN, "Réservation bien annulée.", END)
	} else if ret == 0 {
		fmt.Println(RED, "Erreur : Réservation inexistante", END)
	} else {
		fmt.Println(RED, "Erreur lors de la création de la réservation", END)
	}

	return nil
}
func handle_four() {
	fmt.Println("Visualiser les réservations")
	reservations := res.List_reservations()
	if len(reservations) == 0 {
		fmt.Println(RED, NORES, END)
		return
	}
	for id, reservation := range reservations {
		fmt.Printf(RES, id, reservation.Id, reservation.Room_id, reservation.Date_debut, reservation.Date_fin)
	}
}
func handle_five(scanner *bufio.Scanner) error {
	fmt.Println("Visualiser les réservations d'une salle")
	fmt.Println("Voulez-vous filtrer par date ou par salle ?")
	fmt.Println("1. Par date")
	fmt.Println("2. Par salle")
	fmt.Print("Entrez votre choix : ")
	scanner.Scan()
	choix, err := strconv.Atoi(scanner.Text())
	Handle_err(err)
	if choix == 1 {
		fmt.Print("Entrez la date de début de la réservation : ")
		year, month, day, hour, minute, err := Create_date(scanner)
		if err != nil {
			return errors.New("Annulation")
		}
		date1 := year + "-" + month + "-" + day + " " + hour + ":" + minute
		reservations := res.List_reservations_by_date(date1)
		if len(reservations) == 0 {
			fmt.Println(RED, NORES, END)
			return nil
		}
		for id, reservation := range reservations {
			fmt.Printf(RES, id, reservation.Id, reservation.Room_id, reservation.Date_debut, reservation.Date_fin)
		}
	} else if choix == 2 {
		fmt.Print("Entrez le numéro de la salle ou 0 pour annuler : ")
		scanner.Scan()
		id, err := strconv.Atoi(scanner.Text())
		Handle_err(err)
		if id == 0 {
			return errors.New("Annulation")
		}
		if res.Check_salle(id) == 0 {
			fmt.Println(RED, "Erreur : Salle inexistante", END)
			return nil
		}
		reservations := res.List_reservations_by_room(id)
		if len(reservations) == 0 {
			fmt.Println(RED, NORES, END)
			return nil
		}
		for id, reservation := range reservations {
			fmt.Printf(RES, id, reservation.Id, reservation.Room_id, reservation.Date_debut, reservation.Date_fin)
		}
	} else {
		fmt.Println(RED, "Erreur : Choix incorrect", END)
	}
	return nil
}
func handle_six(scanner *bufio.Scanner) {
	fmt.Println("Exporter les réservations")
	fmt.Print("Entrez le format de l'export (json/csv) : ")
	scanner.Scan()
	input := scanner.Text()
	inputLower := strings.ToLower(input)
	if inputLower == "json" {
		res.Export_reserv_to_json("reservations")
	} else if inputLower == "csv" {
		res.Export_reserv_to_csv("reservations")
	} else {
		fmt.Println("Erreur : Format incorrect")
	}
}
func handle_seven(scanner *bufio.Scanner) {
	fmt.Println("Importer des réservations")
	fmt.Print("Entrez le nom du fichier : ")
	scanner.Scan()
	input := scanner.Text()
	parts := strings.Split(input, ".")
	if parts[len(parts)-1] == "json" {
		res.Import_reserv_from_json(input)
	} else if parts[len(parts)-1] == "csv" {
		res.Import_reserv_from_csv(input)
	} else {
		fmt.Println("Erreur : Format de fichier incorrect")
	}
}

// }
