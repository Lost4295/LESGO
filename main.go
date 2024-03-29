package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func showMenu() {
	fmt.Println()
	fmt.Println("1. Lister les salles disponibles")
	fmt.Println("2. Créer une réservation")
	fmt.Println("3. Annuler une réservation")
	fmt.Println("4. Visualiser les réservations")
	fmt.Println("5. Quitter")
	fmt.Println()
}

func main() {
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
		if number > 5 || number < 1 {
			showMenu()
			fmt.Print("Veuillez entrer un nombre entre 1 et 5 : ")
			continue
		}
		break
	}

	fmt.Println(number)

}
