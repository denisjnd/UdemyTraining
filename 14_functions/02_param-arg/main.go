package main

import "fmt"

func main() {
	greet("John")
}

func greet(name string) {
	fmt.Println("Hi ", name)
}

// greet is declared with a parameter
//when calling greet, pass an argument
/*
func [receiver] func-name(params) return type{

}
*/
