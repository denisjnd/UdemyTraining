package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "The credit y"
		fmt.Println(y)
	}
}
