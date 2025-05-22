package main

import "fmt"

func main() {
	nom_utilisateur := "Sir_King_Über"
	for i := 0; i < len(nom_utilisateur); i++ {
		fmt.Print(string(nom_utilisateur[i]), " ")
	}
}
