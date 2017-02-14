package main

import "fmt"

func main() {
	s := greet("Jane ", "Doe")
	fmt.Println(s)

	fmt.Println(greet("Jane ", "Doe"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname) // Sprint = string print (Its printing to a string instead of standard out)
}
