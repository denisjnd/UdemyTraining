package main

import "fmt"

func main() {
	mySlice := []string{"a", "e", "t", "y", "q", "w"}
	fmt.Println(mySlice[2:4]) // slicing a slice - get from position 2 but dont get upto position 4
	fmt.Println(mySlice[3])   // accessing by index
	fmt.Println("mySlice"[1])
}
