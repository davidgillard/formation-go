package main

import "fmt"

func main() {
	logLevel := "отлаживать"
	for index, runeVal := range logLevel {
		fmt.Println(index, string(runeVal))
	}
}
