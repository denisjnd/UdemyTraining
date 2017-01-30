package main

import "fmt"

func main() {
	a := 40 // declare variable a

	fmt.Println(a)
	fmt.Println(&a) // print the memory address of a

	var b *int = &a // create variable b that is of pointer to an int and assign the value of memory address of a to it
	fmt.Println(b)
}
