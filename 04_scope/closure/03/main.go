package main

import "fmt"

func main() {
	x := 0
	increment := func() int { // anonymous function declaration also called a func expression
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
}

/*
Closure helps us limit the scope of variable used in multiple functions, without closure,
for two or more func to have access to the same variable, that variable will need to be package scope

anonymous function = a function without a name /// Used to put a function inside a function

func expression = assigning a function to a variable
*/
