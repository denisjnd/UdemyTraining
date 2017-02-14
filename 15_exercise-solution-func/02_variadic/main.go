package main

import "fmt"

func main() {
	greatest :=  vari(2,3,4,5,6)
	fmt.Println(greatest)
}

func vari(x ...int) int {
	var largest int
	for _, v := range x {
		if v > largest {
			largest = v
		}
	}
	return largest
}
