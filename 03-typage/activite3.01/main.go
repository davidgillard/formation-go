package main

import "fmt"

func main() {
	coutTotal := .0
	coutTotal += taxeSurVente(0.99, 0.075)
	coutTotal += taxeSurVente(2.75, 0.015)
	coutTotal += taxeSurVente(0.87, 0.02)
	fmt.Println("Total TTC: ", coutTotal)
}

func taxeSurVente(cout float64, taux_taxe_vente float64) float64 {
	return cout * taux_taxe_vente
}
