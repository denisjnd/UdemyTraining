package main

import "fmt"

func main() {
	// call a function with the same function and once you hit return in any function, the code bellow the return
	// will not run again
	fmt.Println(factorial(4))
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
