package main

import "fmt"

type Person struct {
	First string
	Age   int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}

func (p Person) Greeting() {
	fmt.Println("I am just a regular person")
}

// Method overriding in motion ... Use function receivers for method overriding in go lang
func (dz DoubleZero) Greeting() {
	fmt.Println("Miss. MoneyPenny, so good to see you")
}

func main() {
	p1 := Person{
		First: "Denis",
		Age:   30,
	}

	p2 := DoubleZero{
		Person: Person{
			First: "Denison",
			Age:   33,
		},
		LicenseToKill: true,
	}

	p1.Greeting()
	p2.Greeting()
	p2.Person.Greeting()
}
