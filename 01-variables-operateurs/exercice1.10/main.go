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
	name := "David"
	name += " Gillard"
	fmt.Println("Привет,", name)
}
