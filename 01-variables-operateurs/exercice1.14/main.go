package main

import (
	"fmt"
	"time"
)

func main() {
	var compteur1 *int
	compteur2 := new(int)
	compteurTemp := 5
	compteur3 := &compteurTemp
	t := &time.Time{}

	if compteur1 != nil {
		fmt.Printf("Compteur1: %#v\n", *compteur1)
	}
	if compteur2 != nil {
		fmt.Printf("Compteur2: %#v\n", *compteur2)
	}
	if compteur3 != nil {
		fmt.Printf("Compteur3: %#v\n", *compteur3)
	}

	if t != nil {
		fmt.Printf("time: %#v\n", *t)
		fmt.Printf("time: %#v\n", t.String())
	}
}
