package main

import "fmt"

func filtter(numbers []int, callback func(int) bool) []int {
	//xs := []int{}
	var xs []int
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}

	return xs
}

func main() {
	xs := filtter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})

	fmt.Println(xs)
}
