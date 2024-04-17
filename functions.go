package main

import (
	"LESGO/web"
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"

	"github.com/joho/godotenv"
)

const INFO = "\n%s Pour quitter, entrez '0' %s"

func HandleErr(err error) {
	if err != nil {
		fmt.Println("Erreur : ", err)
	}
}

func Start() {
	var verbose, magenta, noWeb, noCli bool = false, false, false, false
	if slices.Contains(os.Args, "-v") || slices.Contains(os.Args, "--verbose") {
		verbose = true
	}
	if slices.Contains(os.Args, "-m") || slices.Contains(os.Args, "--magenta") {
		magenta = true
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

	val = strconv.FormatBool(magenta)
	check = os.Getenv("MAGENTA")
	if check != val && magenta{
		fmt.Println("Setting MAGENTA to ", val)
		os.Setenv("MAGENTA", val)
	} else {
		magenta, _ = strconv.ParseBool(check)
	}

	val = strconv.FormatBool(noWeb)
	check = os.Getenv("NO_WEB")
	if check != val && noWeb {
		fmt.Println("Setting NO_WEB to ", val)
		os.Setenv("NO_WEB", val)
	} else {
		noWeb, _ = strconv.ParseBool(check)
	}
	// fmt.Println("verbose : ", verbose, " noCli : ", noCli, " magenta : ", magenta, " noWeb : ", noWeb)
	// os.Exit(0)
	if magenta {
		fmt.Print(MAGENTA)
	}
	if !noWeb {
		web.Main()
	}
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

func CreateDate(scanner *bufio.Scanner) (string, string, string, string, string, error) {

	// Année
	var year string
	var intYear int
	good := false
	for !good {
		fmt.Printf("\n%sEntrez l'année de réservation (yyyy):%s", GREEN, END)
		scanner.Scan()
		year = scanner.Text()
		if year == "0" {
			return "0", "0", "0", "0", "0", errors.New("Annulation")
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

	//Mois
	var month string
	var intMonth int
	good = false
	for !good {
		fmt.Printf("\n%sEntrez le mois de réservation (mm):%s", GREEN, END)
		fmt.Printf(INFO, BLUE, END)
		scanner.Scan()
		month = scanner.Text()
		if month == "0" {
			return "0", "0", "0", "0", "0", errors.New("Annulation")
		}
		intMonth, err := strconv.Atoi(month)
		if err != nil {
			fmt.Println(CONVERR, err)
			os.Exit(1)
		}
		if intMonth >= 01 && intMonth <= 12 {
			good = true
			if len(month) == 1 {
				month = "0" + month
			}
		} else {
			fmt.Printf("\n%sLe mois doit être compris entre 1 et 12%s", RED, END)
		}
	}

	//Jour
	var day string
	good = false
	for !good {
		fmt.Printf("\n%sEntrez le jour de réservation (dd):%s", GREEN, END)
		fmt.Printf(INFO, BLUE, END)
		scanner.Scan()
		day = scanner.Text()
		if day == "0" {
			return "0", "0", "0", "0", "0", errors.New("Annulation")
		}
		intDay, err := strconv.Atoi(day)
		if err != nil {
			fmt.Println(CONVERR, err)
			os.Exit(1)
		}
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
		if intDay >= 1 && intDay <= 31 {
			good = true
			if len(day) == 1 {
				day = "0" + day
			}
		} else {
			fmt.Printf("%sLe jour doit être compris entre 1 et 31%s", RED, END)
		}
	}

	//Heure
	var hour string
	good = false
	for !good {
		fmt.Printf("\n%sEntrez l'heure de réservation (hh):%s", GREEN, END)
		fmt.Printf(INFO, BLUE, END)
		scanner.Scan()
		hour = scanner.Text()
		if hour == "0" {
			return "0", "0", "0", "0", "0", errors.New("Annulation")
		}
		intHour, err := strconv.Atoi(hour)
		if err != nil {
			fmt.Println(CONVERR, err)
			os.Exit(1)
		}
		if intHour >= 0 && intHour <= 23 {
			good = true
			if len(hour) == 1 {
				hour = "0" + hour
			}
		} else {
			fmt.Printf("\n%sL'heure doit être comprise entre 0 et 23%s", RED, END)
		}
	}

	//Minute
	var minute string
	good = false
	for !good {
		fmt.Printf("\n%sEntrez les minutes de réservation (mm):%s", GREEN, END)
		fmt.Printf(INFO, BLUE, END)
		scanner.Scan()
		minute = scanner.Text()
		if minute == "0" {
			return "0", "0", "0", "0", "0", errors.New("Annulation")
		}
		intMinute, err := strconv.Atoi(minute)
		if err != nil {
			fmt.Println(CONVERR, err)
			os.Exit(1)
		}
		if intMinute >= 0 && intMinute <= 59 {
			good = true
			if len(minute) == 1 {
				minute = "0" + minute
			}
		} else {
			fmt.Printf("%sLes minutes doivent être comprises entre 0 et 59%s", RED, END)
		}
	}
	return year, month, day, hour, minute, nil
}
