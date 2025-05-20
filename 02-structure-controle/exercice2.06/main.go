package main

import (
	"fmt"
	"time"
)

func main() {
	jourNaissance := time.Tuesday
	switch jourNaissance {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("né un jour de la semaine")
	case time.Saturday, time.Sunday:
		fmt.Println("né un weekend")
	default:
		fmt.Println("Erreur, le jour de naissance n'est pas valide")
	}
}
