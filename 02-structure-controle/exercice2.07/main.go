package main

import (
	"fmt"
	"time"
)

func main() {
	switch jourNaissance := time.Tuesday; {
	case jourNaissance == time.Sunday || jourNaissance == time.Saturday:
		fmt.Println("né un weekend")
	default:
		fmt.Println("Né un autre jour")
	}
}
