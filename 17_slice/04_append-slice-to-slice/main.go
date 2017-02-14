package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6}
	otherSlice := []int{7, 8, 9, 10}

	mySlice = append(mySlice, otherSlice...) // appending a slice to a slice

	fmt.Println(mySlice)
}
