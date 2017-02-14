package main

import "fmt"

func main() {
	greeting := func() { // anonymous function // a waay of having a function inside another function
		fmt.Println("Hello...")
	}

	greeting()
}
// anonymous functions don't have names