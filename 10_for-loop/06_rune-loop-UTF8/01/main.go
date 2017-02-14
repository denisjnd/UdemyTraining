package main

import "fmt"

func main() {
	// Conversion ( Casting )
	//fmt.Println([]byte("Hello")) // print the numeric decimal code point
	// Print out the numeric decimal code point of the folloing
	for i := 50; i <= 140; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
}
