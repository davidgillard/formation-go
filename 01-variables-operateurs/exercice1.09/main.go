package main

import (
	"fmt"
)

func main() {
	// course
	var total float64 = 2 * 13
	fmt.Println("Sous-total :", total)

	// boissons
	total = total + (4 * 2.25)
	fmt.Println("Sous-total :", total)

	// reduction de 5 francs
	total = total - 5
	fmt.Println("Sous-total :", total)

	// donner un pourboire de 10%
	pourboire := total * 0.1
	fmt.Println("Pourboire :", pourboire)

	// Découpe de la facture
	decoupe := total / 2
	fmt.Println("Découpe: ", decoupe)

	// Bonus toutes le 5 visites
	nombre_visite := 24
	nombre_visite = nombre_visite + 1

	// Ici, nous allons calculer si le client obtient une récompense. Nous allons d'abord définir le nombre de visites, puis ajouter 1 Francs à cette visite
	reste := nombre_visite % 5

	// Le client obtient une récompense toutes les cinq visites. Si le reste est de 0, il s'agit de l'une de ces visites. Utilisez l'opérateur == pour vérifier si le reste est de 0
	if reste == 0 {
		// Si c'est le cas, affichez un message indiquant qu'il obtient une récompense
		fmt.Println("Avec cette visite, vous avez gagné une récompense.")
	}
}
