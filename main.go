package main

import (
	// "LESGO/web"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"LESGO/db"
	res "LESGO/reservations"
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
	fmt.Println("6. Quitter")
	fmt.Println()
}

func main() {

	// web.Main()
	db.CreateTest()
	// fmt.Println(time.DateTime)
	scanner := bufio.NewScanner(os.Stdin)
	var number int
	var err error

	fmt.Println("Bienvenue dans le Service de Réservation en Ligne")
	fmt.Println("-------------------------------------------------")
	showMenu()
	fmt.Print("Sélectionnez une option : ")

	for {
		scanner.Scan()
		number, err = strconv.Atoi(scanner.Text())
		if err != nil {
			showMenu()
			fmt.Print("Veuillez entrer un nombre valide : ")
			continue
		}
		if number > 6 || number < 1 {
			showMenu()
			fmt.Print("Veuillez entrer un nombre entre 1 et 6 : ")
			continue
		}
		break
	}
	
	fmt.Println("Option choisie :", number)
	switch number {
	case 1:
		fmt.Println("Liste des salles :")
		res.ListRooms()
	case 2: 
		fmt.Println("Entrez la date de la réservation sous le format yyyy-mm-dd hh:min :")
		var date string
		fmt.Scan(&date)
		fmt.Println("Liste des salles disponibles :")
		res.IsFree(date)
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
		fmt.Println("Quitter")
		os.Exit(0)
	}
}
