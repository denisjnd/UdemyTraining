package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "good morning",
		1: "good afternoon",
		2: "good evening",
	}

	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
}
