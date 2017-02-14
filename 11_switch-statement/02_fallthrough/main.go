package main

import "fmt"

func main() {
	switch "Denis" {
	case "Deni":
		fmt.Println("Hi Deni")
	case "Denis":
		fmt.Println("Hi Denis")
		fallthrough
	case "Dei":
		fmt.Println("Hi Dei")
	case "enis":
		fmt.Println("Hi denis 2")
		fallthrough
	case "juan":
		fmt.Println("Hi juan")
	case "Doris":
		fmt.Println("Hi Doris")

	}
}

/*
with fallthrough, you can run more than one statements that are true.
when one is true, fallthrough automatically says the next one is also true
*/
