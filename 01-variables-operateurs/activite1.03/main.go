package main

import "fmt"

func main() {
	compteur := 5
	var message string
	if compteur > 5 {
		message = "Plus grand que 5"
	} else {
		message = "Pas plus grand que 5"
	}
	fmt.Println(message)
}
