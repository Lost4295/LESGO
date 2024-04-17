package main

import (
	"LESGO/db"
	res "LESGO/reservations"
	"LESGO/web"
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
	ROUGE      = "\033[31;01;51m"
	GREEN      = "\033[32;01m"
	BLANC      = "\033[37;07m"
	BLUE	   = "\033[34;01m"
	MAGENTA   = "\033[35;01m"
)

func main() {
	var verbose, magenta bool = false, false
	if slices.Contains(os.Args,"-v") || slices.Contains(os.Args,"--verbose") {
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
	web.Main()
	db.ConnectToDatabase()
	// fmt.Println(time.DateTime)
	scanner := bufio.NewScanner(os.Stdin)
	var number int
	var err error

	for {
		fmt.Printf("%sBienvenue dans le Service de Réservation en Ligne%s", BLUE, END)
		fmt.Println("%s-------------------------------------------------%s", BLANC, END)
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
			// Année
			var year string
			var intYear int
			good := false
			for !good {
				fmt.Println("Entrez l'année de réservation (yyyy):")
				scanner.Scan()
				year = scanner.Text()
				if len(year) == 4 {
					if year[0] == '2' && year[1] == '0' {
						good = true
					} else {
						fmt.Println("L'année doit commencer par 20")
					}
				} else {
					fmt.Println("L'année doit être au format yyyy")
				}
			}

			//Mois
			var month string
			var intMonth int
			good = false
			for !good {
				fmt.Println("Entrez le mois de réservation (mm):")
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
					fmt.Println("Le mois doit être compris entre 1 et 12")
				}
			}

			//Jour
			var day string
			good = false
			for !good {
				fmt.Println("Entrez le jour de réservation (dd):")
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
						fmt.Println("Le jour doit être compris entre 1 et 31")
					}
				case 4, 6, 9, 11:
					if intDay >= 1 && intDay <= 30 {
						good = true
					} else {
						fmt.Println("Le jour doit être compris entre 1 et 30")
					}
				case 2:
					if intYear%4 == 0 && (intYear%100 != 0 || intYear%400 == 0) {
						if intDay >= 1 && intDay <= 29 {
							good = true
						} else {
							fmt.Println("Le jour doit être compris entre 1 et 29")
						}
					} else {
						if intDay >= 1 && intDay <= 28 {
							good = true
						} else {
							fmt.Println("Le jour doit être compris entre 1 et 28")
						}
					}
				}
				if intDay >= 1 && intDay <= 31 {
					good = true
					if len(day) == 1 {
						day = "0" + day
					}
				} else {
					fmt.Println("Le jour doit être compris entre 1 et 31")
				}
			}

			//Heure
			var hour string
			good = false
			for !good {
				fmt.Println("Entrez l'heure de réservation (hh):")
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
					fmt.Println("L'heure doit être comprise entre 0 et 23")
				}
			}

			//Minute
			var minute string
			good = false
			for !good {
				fmt.Println("Entrez les minutes de réservation (mm):")
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
					fmt.Println("Les minutes doivent être comprises entre 0 et 59")
				}
			}

			fmt.Println("Liste des salles disponibles :")
			res.AreFree(year + "-" + month + "-" + day + " " + hour + ":" + minute)
		case 3:
			fmt.Println("Créer une réservation")
			fmt.Print("Entrez le numéro de la salle : ")
			scanner.Scan()
			salle, err := strconv.Atoi(scanner.Text())
			handleErr(err)
			fmt.Print("Entrez la date de la réservation : ")

			// Année
			var year string
			var intYear int
			good := false
			for !good {
				fmt.Println("Entrez l'année de réservation (yyyy):")
				scanner.Scan()
				year = scanner.Text()
				if len(year) == 4 {
					if year[0] == '2' && year[1] == '0' {
						good = true
					} else {
						fmt.Println("L'année doit commencer par 20")
					}
				} else {
					fmt.Println("L'année doit être au format yyyy")
				}
			}

			//Mois
			var month string
			var intMonth int
			good = false
			for !good {
				fmt.Println("Entrez le mois de réservation (mm):")
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
					fmt.Println("Le mois doit être compris entre 1 et 12")
				}
			}

			//Jour
			var day string
			good = false
			for !good {
				fmt.Println("Entrez le jour de réservation (dd):")
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
						fmt.Println("Le jour doit être compris entre 1 et 31")
					}
				case 4, 6, 9, 11:
					if intDay >= 1 && intDay <= 30 {
						good = true
					} else {
						fmt.Println("Le jour doit être compris entre 1 et 30")
					}
				case 2:
					if intYear%4 == 0 && (intYear%100 != 0 || intYear%400 == 0) {
						if intDay >= 1 && intDay <= 29 {
							good = true
						} else {
							fmt.Println("Le jour doit être compris entre 1 et 29")
						}
					} else {
						if intDay >= 1 && intDay <= 28 {
							good = true
						} else {
							fmt.Println("Le jour doit être compris entre 1 et 28")
						}
					}
				}
				if intDay >= 1 && intDay <= 31 {
					good = true
					if len(day) == 1 {
						day = "0" + day
					}
				} else {
					fmt.Println("Le jour doit être compris entre 1 et 31")
				}
			}

			//Heure
			var hour string
			good = false
			for !good {
				fmt.Println("Entrez l'heure de réservation (hh):")
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
					fmt.Println("L'heure doit être comprise entre 0 et 23")
				}
			}

			//Minute
			var minute string
			good = false
			for !good {
				fmt.Println("Entrez les minutes de réservation (mm):")
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
					fmt.Println("Les minutes doivent être comprises entre 0 et 59")
				}
			}

			res.CreateReservation(salle, year+"-"+month+"-"+day+" "+hour+":"+minute)
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
