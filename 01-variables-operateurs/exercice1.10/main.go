package main

import "fmt"

func main() {
	count := 5
	count += 5
	fmt.Println(count)
	count++
	fmt.Println(count)
	count--
	fmt.Println(count)
	count -= 5
	fmt.Println(count)
	nom := "David"
	nom += " Gillard"
	fmt.Println("Привет,", nom)
}
