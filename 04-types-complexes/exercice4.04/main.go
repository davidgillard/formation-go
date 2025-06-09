package main

import "fmt"

func message() string {
	tab := [...]string{
		"prêt",
		"à",
		"Soyez",
		"partir",
	}
	return fmt.Sprintln(tab[2], tab[0], tab[1], tab[3])
}

func main() {
	fmt.Print(message())
}
