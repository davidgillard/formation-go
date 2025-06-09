package main

import "fmt"

func alimTab(tab [10]int) [10]int {
	for i := 0; i < len(tab); i++ {
		tab[i] = i + 1
	}
	return tab
}

func multTab(tab [10]int) [10]int {
	for i := 0; i < len(tab); i++ {
		tab[i] = tab[i] * tab[i]
	}
	return tab
}

func main() {
	var tab [10]int
	fmt.Println("valeur initiale", tab)
	tab = alimTab(tab)
	fmt.Println("valeur modifié par la fonction alimTab", tab)
	tab = multTab(tab)
	fmt.Println("valeur modifié par la fonction multTab", tab)
}
