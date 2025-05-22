package main

import "fmt"

func main() {
	nom_utilisateur := "Sir_King_Über"
	fmt.Println("Bytes: ", len(nom_utilisateur))
	fmt.Println("Runes: ", len([]rune(nom_utilisateur)))
	fmt.Println(string(nom_utilisateur[:10]))
	fmt.Println(string([]rune(nom_utilisateur)[:10]))
}
