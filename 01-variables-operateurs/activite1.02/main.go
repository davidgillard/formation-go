package main

import (
	"fmt"
)

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	a, b := 5, 10
	// appeller la fonction swap ici
	swap(&a, &b)
	fmt.Println(a == 10, b == 5)
}
