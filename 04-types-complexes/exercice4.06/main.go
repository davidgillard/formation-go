package main

import "fmt"

func message() string {
	valeur := ""
	tab := [...]int{1, 2, 3, 4}

	for i := 0; i < len(tab); i++ {
		tab[i] = tab[i] * tab[i]
		valeur += fmt.Sprintf("%v: %v\n", i, tab[i])
	}
	return valeur
}

func main() {
	fmt.Print(message())
}
