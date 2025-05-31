package main

import "fmt"

func CompareTableau() (bool, bool, bool) {
	var tab1 [5]int
	tab2 := [5]int{0}
	tab3 := [...]int{0, 0, 0, 0, 0}
	tab4 := [5]int{0, 0, 0, 0, 9}

	return tab1 == tab2, tab1 == tab3, tab1 == tab4
}

func main() {
	comp1, comp2, comp3 := CompareTableau()
	fmt.Println("[5]int == [5]int{0}     :", comp1)
	fmt.Println("[5]int == [...]int{0, 0, 0, 0, 0}:", comp2)
	fmt.Println("[5]int == [5]int{0, 0, 0, 0, 9} :", comp3)
}
