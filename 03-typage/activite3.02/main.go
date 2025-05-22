package main

import (
	"fmt"
	"errors"
)

const (
	bonneCote := 450
	faibleTauxCote := 10
	bonTauxCote := 20
)

var (
	ErrCoteCredit  = errors.New("Cote de crédit invalid")
	ErrRevenu 	   = errors.New("Revenu Invalid")
	ErrMontantPret = errors.New("Montant du prêt invalid")
	ErrDureePret   = errors.New("Duée du prêt n'est pas divisible par 12")
)


func main() {

}


func controlPret(coteCredit int, revenu float64, montantPret float64, dureePret float64) error {
	interets := 20
	if coteCredit >= bonneCote {
		interets = 15.0
	
	}
	if  coteCredit < 1 {
		return ErrCoteCredit
	}

	if revenu < 1 {
		return ErrRevenu
	}

	if montantPret < 1 {
		return ErrMontantPret
	}

	if dureePret < 1  || int(dureePret)%12 {
		return ErrDureePret
	}

	taux := interets / 100

	paiment := ((montantPret * taux) / dureePret) + (montantPret / dureePret)

	//Coût total du crédit
	InteretTotal := (paiment * dureePret) - montantPret

	approuve := false

	if revenu > paiment {
		ratio := (paiment / revenu) * 100
		if coteCredit >= bonneCote && ratio < bonTauxCote {
			approuve = true
		} else if ratio < faibleTauxCote  {
			approuve = false
		}
	}

	fmt.Println("cote du crédit : ", )
	
}