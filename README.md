# LESGO

Bienvenue sur notre plateforme de réservation en ligne ! Simplifiez vos réservations pour voyages, événements et services avec notre interface conviviale et nos fonctionnalités intuitives. Profitez d'un processus fluide, sécurisé et rapide pour concrétiser vos projets en quelques clics seulement.
## Fonctionnalités

- Interface Utilisateur en CLI
- Gestion de réservations
- Persistence des données
- Export des réservations
- Interface Web


## Installation

L'application est écrite en GO. Elle sera exécutée avec le programme de compilation Go. Il est nécessaire d'installer le binaire avant de pouvoir installer le projet. 

- Installer Go

[Lien vers la page de téléchargement](https://go.dev/dl/)

- Cloner le repository dans l'endroit de votre choix

```bash
  git clone https://github.com/Lost4295/LESGO.git
  cd LESGO
```
- Modifier le fichier .env à votre convenance

Il est nécessaire que la base de données soit créée au préalable avant de pouvoir lancer l'application. Assurez-vous que l'utilisateur utilisé a le droit de lire et écrire dans la base de données sélectionnée.

> [!WARNING]
> Si vous changez les variables d'environnement `PORT`, `USER`, `PASSWORD` et `DBNAME`, assurez-vous que les données sont bien conformes aux données entrées dans le fichier docker-compose.yml. 

## Variables d'Environnement 

Pour lancer ce projet, il faut que dans le fichier .env, il y ait les variables suivantes :

`VERBOSE` : `"true"` ou `"false"`

`NO_CLI` : `"true"` ou `"false"`

`CLEAR` : `"true"` ou `"false"`

`NO_WEB` : `"true"` ou `"false"`

`HOST` : l'hôte de la base de données

`PORT` : le port utilisé pour se connecter à la base de données

`DBNAME` : le nom de la base de données

`USER` : l'utilisateur de la base de données

`PASSWORD` : le mot de passe de la base de données

`WEB_PORT` : le port utilisé pour utiliser le serveur web


Toutes les variables doivent être des chaines de caractères entourées de \" (double quotes).

## Lancement

Pour lancer le programme, il faut lancer le programme go. 

```bash
  go run .
```

Si vous nécessitez le besoin de dockeriser la connexion à votre base de données, veuillez lancer la commande suivante également : 

```bash
  docker-compose up
```



## Auteurs

- [@Lost4295](https://www.github.com/Lost4295)
- [@LumineWollah](https://www.github.com/LumineWollah)
- [@Leopold194](https://www.github.com/Leopold194)

