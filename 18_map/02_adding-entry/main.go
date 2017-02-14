package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"Tim" : "Good morning",
		"Joe" : "Good afternoon",
	}

	myGreeting["Helen"] = "Helena"

	fmt.Println(len(myGreeting)) // get the length of the map
}
