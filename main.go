package main

import (
	res "LESGO/reservations"
	"bufio"
	"fmt"
	"os"
	"slices"
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
	WHITEONRED = "\033[37;41m"
	END        = "\033[0m"
	RED        = "\033[31;01;51m"
	GREEN      = "\033[32;01m"
	WHITE      = "\033[37;07m"
	BLUE       = "\033[34;01m"
	MAGENTA    = "\033[35;01m"
)

func main() {
	var verbose, magenta bool = false, false
	if slices.Contains(os.Args, "-v") || slices.Contains(os.Args, "--verbose") {
		verbose = true
	}

	if slices.Contains(os.Args, "-m") || slices.Contains(os.Args, "--magenta") {
		magenta = true
	}

	os.Setenv("verbose", strconv.FormatBool(verbose))
	os.Setenv("magenta", strconv.FormatBool(magenta))
	if magenta {
		fmt.Print(MAGENTA)
	}
	// web.Main()
	// db.ConnectToDatabase()
	// fmt.Println(time.DateTime)
	scanner := bufio.NewScanner(os.Stdin)
	var number int
	var err error

	for {
		fmt.Printf("%sBienvenue dans le Service de Réservation en Ligne%s", BLUE, END)
		fmt.Printf("\n%s-------------------------------------------------%s", WHITE, END)
		showMenu()
		fmt.Print("%sSélectionnez une option : %s", GREEN, END)

		scanner.Scan()
		number, err = strconv.Atoi(scanner.Text())
		if err != nil {
			showMenu()
			fmt.Print("%sVeuillez entrer un nombre valide : %s", RED, END)
			continue
		}
		if number > 8 || number < 1 {
			showMenu()
			fmt.Print("%sVeuillez entrer un nombre entre 1 et 8 : %s", RED, END)
			continue
		}

		fmt.Println("Option choisie :", number)
		switch number {
		case 1:
			fmt.Printf("\n%sListe des salles :%s", BLUE, END)
			res.ListRooms()
		case 2:
			// Année
			var year string
			var intYear int
			good := false
			for !good {
				fmt.Printf("\n%sEntrez l'année de réservation (yyyy):%s", GREEN, END)
				scanner.Scan()
				year = scanner.Text()
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
				scanner.Scan()
				month = scanner.Text()
				intMonth, err := strconv.Atoi(month)
				if err != nil {
					fmt.Println("Erreur de conversion: ", err)
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
				scanner.Scan()
				day = scanner.Text()
				intDay, err := strconv.Atoi(day)
				if err != nil {
					fmt.Println("Erreur de conversion: ", err)
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
					fmt.Printf("\n%sLe jour doit être compris entre 1 et 31%s", RED, END)
				}
			}

			//Heure
			var hour string
			good = false
			for !good {
				fmt.Printf("\n%sEntrez l'heure de réservation (hh):%s", GREEN, END)
				scanner.Scan()
				hour = scanner.Text()
				intHour, err := strconv.Atoi(hour)
				if err != nil {
					fmt.Println("Erreur de conversion: ", err)
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
				scanner.Scan()
				minute = scanner.Text()
				intMinute, err := strconv.Atoi(minute)
				if err != nil {
					fmt.Println("Erreur de conversion: ", err)
					os.Exit(1)
				}
				if intMinute >= 0 && intMinute <= 59 {
					good = true
					if len(minute) == 1 {
						minute = "0" + minute
					}
				} else {
					fmt.Printf("\n%sLes minutes doivent être comprises entre 0 et 59%s", RED, END)
				}
			}

			fmt.Printf("\n%sListe des salles disponibles :%s", BLUE, END)
			res.AreFree(year + "-" + month + "-" + day + " " + hour + ":" + minute)
		case 3:
			fmt.Printf("\n%sCréer une réservation%s", BLUE, END)
			fmt.Print("%sEntrez le numéro de la salle : %s", GREEN, END)
			scanner.Scan()
			salle, err := strconv.Atoi(scanner.Text())
			handleErr(err)

			// Année
			var year string
			var intYear int
			good := false
			for !good {
				fmt.Printf("\n%sEntrez l'année de réservation (yyyy):%s", GREEN, END)
				scanner.Scan()
				year = scanner.Text()
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
				scanner.Scan()
				month = scanner.Text()
				intMonth, err := strconv.Atoi(month)
				if err != nil {
					fmt.Println("Erreur de conversion: ", err)
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
				scanner.Scan()
				day = scanner.Text()
				intDay, err := strconv.Atoi(day)
				if err != nil {
					fmt.Println("Erreur de conversion: ", err)
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
					fmt.Println("Le jour doit être compris entre 1 et 31%s", RED, END)
				}
			}

			//Heure
			var hour string
			good = false
			for !good {
				fmt.Printf("\n%s%sEntrez l'heure de réservation (hh):%s", GREEN, END)
				scanner.Scan()
				hour = scanner.Text()
				intHour, err := strconv.Atoi(hour)
				if err != nil {
					fmt.Println("Erreur de conversion: ", err)
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
				scanner.Scan()
				minute = scanner.Text()
				intMinute, err := strconv.Atoi(minute)
				if err != nil {
					fmt.Println("Erreur de conversion: ", err)
					os.Exit(1)
				}
				if intMinute >= 0 && intMinute <= 59 {
					good = true
					if len(minute) == 1 {
						minute = "0" + minute
					}
				} else {
					fmt.Printf("\n%sLes minutes doivent être comprises entre 0 et 59%s", RED, END)
				}
			}

			res.CreateReservation(salle, year+"-"+month+"-"+day+" "+hour+":"+minute)
		case 4:
			fmt.Printf("\n%sAnnuler une réservation%s", BLUE, END)
			fmt.Print("%sEntrez le numéro de la réservation : %s", GREEN, END)
			scanner.Scan()
			id, err := strconv.Atoi(scanner.Text())
			handleErr(err)
			res.DeleteReservation(id)
		case 5:
			fmt.Printf("\n%sVisualiser les réservations%s", BLUE, END)
			res.ListReservations()
		case 6:
			fmt.Printf("\n%sExporter les réservations%s", BLUE, END)
			fmt.Print("%sEntrez le format de l'export (json/csv) : %s", GREEN, END)
			scanner.Scan()
			input := scanner.Text()
			inputLower := strings.ToLower(input)

			if inputLower == "json" {
				res.ExportReservToJson("reservations")
			} else if inputLower == "csv" {
				res.ExportReservToCSV("reservations")
			} else {
				fmt.Printf("\n%sErreur : Format incorrect%s", RED, END)
			}
		case 7:
			fmt.Printf("\n%sImporter des réservations%s", BLUE, END)
			fmt.Print("%sEntrez le nom du fichier : %s", GREEN, END)
			scanner.Scan()
			input := scanner.Text()
			parts := strings.Split(input, ".")

			if parts[len(parts)-1] == "json" {
				res.ImportReservFromJson(input)
			} else if parts[len(parts)-1] == "csv" {
				res.ImportReservFromCSV(input)
			} else {
				fmt.Printf("\n%sErreur : Format de fichier incorrect%s", RED, END)
			}
		case 8:
			fmt.Printf("\n%sQuitter%s", BLUE, END)
			os.Exit(0)
		}
		fmt.Printf("\n%sAppuyer sur n'importe quelle touche pour revenir au menu principal%s", GREEN, END)
		scanner.Scan()
	}
}
