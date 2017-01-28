package main

import "fmt"

func main() {
	a := 20
	//b := "Denis"
	//c := 4.17
	//d := true

	// multiple assignment
	var b, c string = "store in B", "Store in C"

	fmt.Printf("%v \n", a)
	fmt.Printf("%T \n", a)

	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
}
