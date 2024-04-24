package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	WHITEONRED = "\033[37;41m"
	END        = "\033[0m"
	RED        = "\033[31;01;51m"
	GREEN      = "\033[32;01m"
	BLANC      = "\033[37;07m"
	BLUE       = "\033[34;01m"
	WHITE      = "\033[37;07m"
	RES        = "%d - Reservation n°%d : Salle %d réservée du %s au %s\n"
)

func main() {
	Start()

	// Si vous réussissez à faire ça, allez y
	// Il faut pouvoir lancer le programme sans le CLI
	// noCli, err := strconv.ParseBool(os.Getenv("NO_CLI"))
	// HandleErr(err)
	clear, _ := strconv.ParseBool(os.Getenv("CLEAR"))

	scanner := bufio.NewScanner(os.Stdin)
	var number int = -1
	// if !noCli {
	fmt.Printf("%sBienvenue dans le Service de Réservation en Ligne%s", BLUE, END)
	for number != 0 {

		fmt.Printf("\n%s~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~%s", WHITE, END)
		Show_menu()
		fmt.Printf("%sSélectionnez une option : %s", GREEN, END)
		scanner.Scan()
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			Show_menu()
			number = -1
			fmt.Printf("%sVeuillez entrer un nombre valide : %s", RED, END)
			continue
		}
		if 0 > number || number > 8 {
			Show_menu()
			fmt.Printf("%sVeuillez entrer un nombre entre 0 et 8 : %s", RED, END)
			continue
		}
		fmt.Println("Option choisie :", number)
		if clear {
			Clear()
		}
		switch number {
		case 0:
			handle_zero()
		case 1:
			handle_one()
		case 2:
			err = handle_two(scanner)
			if err != nil {
				fmt.Println(RED, "Annulation", END)
				continue
			}
		case 3:
			err = handle_three(scanner)
			if err != nil {
				fmt.Println(RED, "Annulation", END)
				continue
			}
		case 4:
			handle_four()
		case 5:
			err = handle_five(scanner)
			if err != nil {
				fmt.Println(RED, "Annulation", END)
				continue
			}
		case 6:
			handle_six(scanner)
		case 7:
			handle_seven(scanner)
		}
	}
}
