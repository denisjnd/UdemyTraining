package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 40
	fmt.Printf("%T %v\n", myAge, myAge)

	var yourAge int
	yourAge = 45
	fmt.Printf("%T %v\n\n", yourAge, yourAge)

	// this will not work cos there are of different types
	//fmt.Println(myAge + yourAge)

	// but using conversion will work
	fmt.Println(int(myAge) + yourAge)
}
