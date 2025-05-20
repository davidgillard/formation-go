package main

import "fmt"

func main() {
	valeur := 4
	if valeur < 0 {
		fmt.Println("la valeur saisie ne peut être négative")
	} else if valeur%2 == 0 {
		fmt.Println(valeur, "est pair")
	} else {
		fmt.Println(valeur, "est impair")
	}
}
