package main

import "fmt"

func ajout5Value(compteur int) {
	compteur += 5
	fmt.Println("ajout5Value :", compteur)
}

func ajout5Point(compteur *int) {
	*compteur += 5
	fmt.Println("ajout5Point :", *compteur)
}

func main() {
	var compteur int
	ajout5Value(compteur)
	fmt.Println("ajout5Value post:", compteur)
	ajout5Point(&compteur)
	fmt.Println("ajout5Point post:", compteur)
}
