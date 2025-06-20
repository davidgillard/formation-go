package main

import "fmt"

func getTab() [10]int {
	var tab [10]int
	for i := 0; i < 10; i++ {
		tab[i] = i + 1
	}
	return tab
}

func main() {
	fmt.Println(getTab())
}
