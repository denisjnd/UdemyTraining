package main

import "fmt"

func main() {
	switch "Denis" {
	case "Deni":
		fmt.Println("Hi Deni")
	case "Denis":
		fmt.Println("Hi Denis")
	default:
		fmt.Println("You have no friends")

	}
}
