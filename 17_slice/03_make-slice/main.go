package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 5) // make(slice, length, capacity).. make([]int, 3) // 3 is both length and capacity

	fmt.Println("------------------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("------------------------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len: ", len(mySlice), "Cap: ", cap(mySlice), "Value: ", mySlice[i])
	}
}
