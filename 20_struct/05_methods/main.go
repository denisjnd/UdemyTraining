package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() string { // the receiver attaches the function to the person type
	return p.first + p.last
}

func main() {
	p1 := person{"Denis", "John", 30}
	p2 := person{"Denison", "Johnson", 40}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}
