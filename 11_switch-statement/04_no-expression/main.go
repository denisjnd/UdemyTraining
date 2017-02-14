package main

import "fmt"

func main() {
	name := "Juliet"

	switch {
	case len(name) == 6:
		fmt.Println("Name is length 6")
	case name == "Juliet":
		fmt.Println("Hi Juliet")
	case name == "Juan", name == "John":
		fmt.Println("Hi juan or John")
	default:
		fmt.Println("No match found")

	}
}
