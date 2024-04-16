package db

import (
	"log"
	"math/rand"
)

// Connect creates a connection to the database

func ConnectToDatabase() {
	db, err := Connect("user", "password")
	if err != nil {
		log.Fatalln("Une erreur est survenue ! \n", err, "Essayez de voir si la base de données est bien accessible en vérifiant que le conteneur docker est bien lancé avec 'docker compose up'.")
	}
	defer db.Close()
	db, err = initDB(db)
	if err != nil {
		log.Fatalln("Une erreur est survenue ! \n\n", err, "\n\nEssayez de voir si la base de données est bien accessible en vérifiant que le conteneur docker est bien lancé avec 'docker compose up'.")
	}
	for i := 0; i < 10; i++ {
		addRoom(db, "Room", rand.Intn(100))
	}
}
