package main

import "fmt"

func main() {
	if false {
		fmt.Println("First statement")
	} else if true {
		fmt.Println("middle statement")
	} else {
		fmt.Println("Second Statement")
	}
}
