package main

import "fmt"

func main() {
	visites := 15
	fmt.Println("Première visite : ", visites == 1)
	fmt.Println("Visite en retour : ", visites != 1)
	fmt.Println("Membre Argent : ", visites >= 10 && visites < 21)
	fmt.Println("Membre Or : ", visites >= 20 && visites <= 30)
	fmt.Println("Membre Platine : ", visites > 30)
}
