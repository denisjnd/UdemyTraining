package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

type DoubleZero struct {
	Person
	First         string
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "Denis",
			Last:  "John",
			Age:   30,
		},
		First:         "Denison",
		LicenseToKill: true,
	}
	// fields and methods of the inner-types are promoted to the outer-type
	fmt.Println(p1.First)
	fmt.Println(p1.Person.First)
}
