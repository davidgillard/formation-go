package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() { // création de la fonction main
	rand.Seed(time.Now().UnixNano()) // Alimentez le générateur de nombres aléatoires permettant de gerer le nombre d'étoile
	r := rand.Intn(5)                // Générer un nombre aléatoire compris entre 0 et 1 pour obtenir un nombre compris entre 1 et 5
	stars := strings.Repeat("*", r)  // Utiliser le répéteur de chaîne pour créer une chaîne contenant le nombre d'étoiles souhaité
	fmt.Println(stars)               //Afficher la chaîne contenant les étoiles dans la console
}
