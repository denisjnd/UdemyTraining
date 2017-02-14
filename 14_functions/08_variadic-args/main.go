package main

import "fmt"

func main() {
	data := []float64{2,4}
	n := average(data...)
	println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64

	for _, v:= range sf {
		total += v
	}

	return total / float64(len(sf))
}
