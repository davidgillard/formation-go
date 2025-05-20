package main

import "fmt"

func main() {
	valeur := 11
	if valeur%2 == 0 {
		fmt.Println(valeur, "est pair")
	} else {
		fmt.Println(valeur, "est impair")
	}
}
