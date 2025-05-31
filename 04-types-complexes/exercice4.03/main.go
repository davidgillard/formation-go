package main

import "fmt"

func CompareTableau() (bool, bool, [10]int) {
	var tab1 [10]int
	// définir la clé 9 sur la valeur 0
	tab2 := [...]int{9: 0}
	/* définit la clé 0 sur la valeur 1, définir la clé 9 sur la valeur 10
	et enfin définir la clé 4 sur la valeur 5 */
	tab3 := [10]int{1, 9: 10, 4: 5}

	return tab1 == tab2, tab1 == tab3, tab3
}

func main() {
	comp1, comp2, tab3 := CompareTableau()
	fmt.Println("[10]int == [...]{9:0}     :", comp1)
	fmt.Println("[10]int == [10]int{1, 9:10, 4, 5}:", comp2)
	fmt.Println("tab3				:", tab3)
}
