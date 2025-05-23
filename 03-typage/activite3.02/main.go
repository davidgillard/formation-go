package main

import (
	"errors"
	"fmt"
)

const (
	bonneCote      = 450
	faibleTauxCote = 10
	bonTauxCote    = 20
)

var (
	ErrCoteCredit  = errors.New("Cote de crédit invalid")
	ErrRevenu      = errors.New("Revenu Invalid")
	ErrMontantPret = errors.New("Montant du prêt invalid")
	ErrDureePret   = errors.New("Duée du prêt n'est pas divisible par 12")
)

func main() {
	// approuvé
	fmt.Println("Candidat 1")
	fmt.Println("----------")
	err := controlPret(500, 1000, 1000, 24)
	if err != nil {
		fmt.Println("Erreur: ", err)
		return
	}

	// refusé
	fmt.Println("Candidat 2")
	fmt.Println("----------")
	err = controlPret(350, 1000, 10000, 12)
	if err != nil {
		fmt.Println("Erreur: ", err)
		return
	}

	fmt.Println("Candidat 3")
	fmt.Println("----------")
	err = controlPret(350, 3700, 250000, 12)
	if err != nil {
		fmt.Println("Erreur: ", err)
		return
	}

}

func controlPret(coteCredit int, revenu float64, montantPret float64, dureePret float64) error {
	interets := 20.0
	if coteCredit >= bonneCote {
		interets = 15.0

	}
	if coteCredit < 1 {
		return ErrCoteCredit
	}

	if revenu < 1 {
		return ErrRevenu
	}

	if montantPret < 1 {
		return ErrMontantPret
	}

	if dureePret < 1 || int(dureePret)%12 != 0 {
		return ErrDureePret
	}

	taux := interets / 100
	paiment := ((montantPret * taux) / dureePret) + (montantPret / dureePret)

	//Coût total du crédit
	interetTotal := (paiment * dureePret) - montantPret

	approuve := false

	if revenu > paiment {
		ratio := (paiment / revenu) * 100
		if coteCredit >= bonneCote && ratio < bonTauxCote {
			approuve = true
		} else if ratio < faibleTauxCote {
			approuve = false
		}
	}

	fmt.Println("Cote du crédit: ", coteCredit)
	fmt.Println("Revenu: ", revenu)
	fmt.Println("Montant du prêt: ", montantPret)
	fmt.Println("Durée du prêt: ", dureePret)
	fmt.Println("Paiment mensuel: ", paiment)
	fmt.Println("Taux: ", interets)
	fmt.Println("Cout total: ", interetTotal)
	fmt.Println("Cote du crédit : ", coteCredit)
	fmt.Println("Approuvé: ", approuve)
	fmt.Println("")

	return nil
}
