package main

import "fmt"

func main() {
	valeur := 16
	if valeur%2 == 0 {
		fmt.Println(valeur, "est pair")
	}
	if valeur%2 == 1 {
		fmt.Println(valeur, "est impair")
	}
}
