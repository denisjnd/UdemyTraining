package main

import "fmt"

func main() {
	a := 40 // declare variable a

	fmt.Println(a)
	fmt.Println(&a) // print the memory address of a

	var b *int = &a // create variable b that is of pointer to an int and assign the value of memory address of a to it
	fmt.Println(b)

	fmt.Println(*b) // dereference

	/*
		b is an int pointer
		b points to the memory address where an int is stored
		to see the value of the memory address, add * in front of b
		this is known as dereferencing
		the * is an operator in this case
	*/

}
