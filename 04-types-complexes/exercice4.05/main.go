package main

import "fmt"

func message() string {
	tab := [4]string{"prêt", "Soyez", "de", "partir"}
	tab[1] = "C'est"
	tab[0] = "l'heure"
	return fmt.Sprintln(tab[1], tab[0], tab[2], tab[3])
}

func main() {
	fmt.Print(message())
}
