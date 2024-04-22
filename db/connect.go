package db

import (
	"log"
	// "math/rand"
	"os"
)

const (
	WHITEONRED = "\033[37;41m"
	END        = "\033[0m"
	ROUGE      = "\033[31;01;51m"
)

func CheckConnection() {
	/*
		Function allowing you to see if the database is accessible and if the creation went well
	*/
	db, err := Connect(os.Getenv("USER"), os.Getenv("PASSWORD"))
	if err != nil {
		log.Fatalln(WHITEONRED, "Une erreur est survenue !", END, "\n\n", err, ROUGE, "\n\nEssayez de voir si la base de données est bien accessible en vérifiant que le conteneur docker est bien lancé avec 'docker compose up'.", END)
	}
	defer db.Close()
	/*db*/ _, err = initDB(db)
	if err != nil {
		log.Fatalln(WHITEONRED, "Une erreur est survenue !", END, "\n\n", err, ROUGE, "\n\nEssayez de voir si la base de données est bien accessible en vérifiant que le conteneur docker est bien lancé avec 'docker compose up'.", END)
	}
	// To create fake rooms
	// for i := 0; i < 10; i++ {
	// 	addRoom(db, "Room", rand.Intn(100))
	// }
}
