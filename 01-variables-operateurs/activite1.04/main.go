package main

import "fmt"

func main() {
	compteur := 0
	if compteur < 5 {
		compteur = 10
		compteur++
	}
	fmt.Println(compteur == 11)
}
