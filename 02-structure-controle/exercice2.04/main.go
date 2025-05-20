package main

import (
	"errors"
	"fmt"
)

func validation(valeur int) error {
	if valeur < 0 {
		return errors.New("la valeur saisie ne peut être négative")
	} else if valeur > 100 {
		return errors.New("la valeur saisie ne peut pas exceder 100")
	} else if valeur%7 == 0 {
		return errors.New("la valeur saisie n'est pas divisible par 7")
	} else {
		return nil
	}
}

func main() {
	valeur := 21
	if err := validation(valeur); err != nil {
		fmt.Println(err)
	} else if valeur%2 == 0 {
		fmt.Println(valeur, "est pair")
	} else {
		fmt.Println(valeur, "est impair")
	}
}
