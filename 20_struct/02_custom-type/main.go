package main

import "fmt"

type foo int // alias of type int is foo

func main() {
	var myAge foo
	myAge = 40
	fmt.Printf("%T %v \n", myAge, myAge)
}
