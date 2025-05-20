package main

import "fmt"

func main() {

	mots := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}

	motPlusUtilise := ""
	nbOccurence := 0
	for key, value := range mots {
		fmt.Println(key, "=", value)
		if value > nbOccurence {
			nbOccurence = value
			motPlusUtilise = key
		}
	}
	fmt.Println("le mot le plus utilisé: ", motPlusUtilise)
	fmt.Println("avec un nombre de: ", nbOccurence)

}
