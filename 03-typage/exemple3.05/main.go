package main

import "fmt"

func main() {
	commentaire1 := `C'est le MEILLEUR 
	que jamais!`
	commentaire2 := `C'est le MEILLEUR\nque jamais!`
	commentaire3 := "C'est le MEILLEUR\nque jamais!"
	fmt.Print(commentaire1, "\n\n")
	fmt.Print(commentaire2, "\n\n")
	fmt.Print(commentaire3, "\n")
}
