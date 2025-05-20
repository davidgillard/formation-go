package main

import "fmt"

func main() {
	prenoms := []string{"Daniel", "Patricia", "Marco", "Pascal", "Nathalie", "David"}
	for i := 0; i < len(prenoms); i++ {
		fmt.Println(prenoms[i])
	}
}
