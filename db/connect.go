package db

import (
	"log"
	"math/rand"
)
const (
	WHITEONRED = "\033[37;41m"
	END   = "\033[0m"
	ROUGE = "\033[31;01;51m"
	GREEN = "\033[32;01m"
	BLANC = "\033[37;07m"
	BLUE  = "\033[34;01m"
)
// Connect creates a connection to the database

func ConnectToDatabase() {
	db, err := Connect("user", "password")
	if err != nil {
		log.Fatalln(WHITEONRED,"Une erreur est survenue !", END,"\n\n", err, ROUGE,"\n\nEssayez de voir si la base de données est bien accessible en vérifiant que le conteneur docker est bien lancé avec 'docker compose up'.",END)
	}
	defer db.Close()
	db, err = initDB(db)
	if err != nil {
		log.Fatalln(WHITEONRED,"Une erreur est survenue !", END,"\n\n", err, ROUGE,"\n\nEssayez de voir si la base de données est bien accessible en vérifiant que le conteneur docker est bien lancé avec 'docker compose up'.",END)
	}
	for i := 0; i < 10; i++ {
		addRoom(db, "Room", rand.Intn(100))
	}
}
