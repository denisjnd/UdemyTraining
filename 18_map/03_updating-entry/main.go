package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"Tim" : "Good morning",
		"Joe" : "Good afternoon",
	}

	myGreeting["Helen"] = "Helena"

	fmt.Println(myGreeting["Tim"])

	myGreeting["Tim"] = "Good evening" // updating a entry

	fmt.Println("Updated Entry: ",myGreeting["Tim"])
}
