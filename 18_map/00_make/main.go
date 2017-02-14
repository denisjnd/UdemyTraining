package main

import "fmt"

func main() {
	var myGreeting = make(map[string]string) // make(map[var.type for key]var.type for value)
	myGreeting["Tim"] = "Good morning"
	myGreeting["Joe"] = "Good afternoon"

	fmt.Println(myGreeting)
}
