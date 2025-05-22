package main

import "fmt"

func main() {
	nom_utilisateur := "Sir_King_Über"
	runes := []rune(nom_utilisateur)
	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i]), " ")
	}
}
