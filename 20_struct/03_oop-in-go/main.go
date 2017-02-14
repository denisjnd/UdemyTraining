package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

type DoubleZero struct {
	Person        // anonymous type. all the fields in Person struct(type) is visible here ... The inner type gets promoted to the outer type. This is Inheritance(Reusability)
	First         string
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			first: "Denis",
			last:  "John",
			age:   40,
		},
		First:         "Double Zero Seven",
		LicenseToKill: true,
	}

	p2 := DoubleZero{
		Person: Person{
			first: "Miss",
			last:  "Money Penny",
			age:   30,
		},
		First:         "If looks could kill",
		LicenseToKill: false,
	}

	fmt.Println(p1.First, p1.Person.first)
	fmt.Println(p2.first, p2.LicenseToKill)
}
