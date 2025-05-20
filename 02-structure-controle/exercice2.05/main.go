package main

import (
	"fmt"
	"time"
)

func main() {
	jourNaissance := time.Saturday
	switch jourNaissance {
	case time.Monday:
		fmt.Println("L'enfant du lundi a un beau visage")
	case time.Tuesday:
		fmt.Println("L'enfant du mardi est plein de grâce")
	case time.Wednesday:
		fmt.Println("L'enfant du mercredi est plein de malheurs")
	case time.Thursday:
		fmt.Println("L'enfant du jeudi a encore beaucoup à faire")
	case time.Friday:
		fmt.Println("L'enfant du vendredi est aimant et généreux")
	case time.Saturday:
		fmt.Println("L'enfant du samedi travaille dur pour gagner sa vie")
	case time.Sunday:
		fmt.Println("L'enfant du dimanche est beau et joyeux")
	default:
		fmt.Println("Erreur, le jour de naissance n'est pas valide")
	}
}
