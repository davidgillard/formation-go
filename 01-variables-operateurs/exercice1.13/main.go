package main

import (
	"fmt"
	"time"
)

func main() {
	var compte1 *int
	compte2 := new(int)
	compteTemp := 5
	compte3 := &compteTemp
	t := &time.Time{}

	fmt.Printf("compte1 : %#v\n", compte1)
	fmt.Printf("compte2 : %#v\n", compte2)
	fmt.Printf("compte3 : %#v\n", compte3)
	fmt.Printf("time : %#v\n", t)

}
