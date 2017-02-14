package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

// another shape
type Circle struct {
	radius float64
}

// This implements the Shape interface
func (s Square) area() float64 {
	return s.side * s.side
}

// This also implements the Shape interface
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// shape interface that the functions will implement
type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Square{10}
	c := Circle{20}
	info(s)
	info(c)
}
