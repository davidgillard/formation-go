package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	a := make([]int, 20)
	for i := 0; i < 20; i++ {
		a[i] = r.Intn(20)
	}
	fmt.Println("Mon slice non trié: ", a)
	for swap := true; swap; {
		swap = false
		for i := 1; i < len(a); i++ {
			if a[i-1] > a[i] {
				a[i], a[i-1] = a[i-1], a[i]
				swap = true
			}
		}
	}
	fmt.Println("Mon slice trié: ", a)
}
