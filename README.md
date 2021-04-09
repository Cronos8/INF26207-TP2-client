# INF26207-TP2-client

## Pour commencer

Simulation d'un client UDP avec les sockets dans le cadre d'un travail pratique du cours de téléinformatique à l'UQAR.

### Pré-requis

- Go v1.15.6

Testé sous :
- macOS 11.2.3
- windows 10 

### Installation

Dans la racine du répertoire, exécuter la commande : ``go build .``.

## Démarrage

À la racine du répertoire :

Exécuter la commande :``./INF26207-TP2-client IPServeur:PortServeur extensionFichier``
Exemple : ``./INF26207-TP2-client 127.0.0.1:22222 jpeg``

Une fois la transmission de fichier terminé le fichier sera visible à la racine sous le nom de ``packet.extension``.

## Suppression de l'exécutable 

À la racine du répertoire :

Exécuter la commande :``go clean``.

## Auteurs

Alexandre Nguyen