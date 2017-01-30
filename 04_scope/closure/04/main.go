package main

import "fmt"

func wrapper() func() int { // func wrapper is returning a function, and that function is returning an int
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}
