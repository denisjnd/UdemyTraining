package main

import "fmt"

type Square struct {
	size float64
}

func (z Square) area() float64 {
	return z.size * z.size
}

type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Square{10}
	info(s)
}
