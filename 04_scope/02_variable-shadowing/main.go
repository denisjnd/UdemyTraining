package main

import "fmt"

func max(x int) int  {
	return 42 * x;
}

func main() {
	max := max(2);
	fmt.Println(max);
}