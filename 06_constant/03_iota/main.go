package main

import "fmt"

const (
	A = iota // 0
	B        // 1
	C        // 2
)

const (
	_  = iota             // 0 is gone
	KM = 1 << (iota * 10) // (1 * 10)
	MB = 1 << (iota * 10) // (2 * 10)
)

func main() {

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

	fmt.Println("binary\t\t\tdecimal")
	fmt.Println("%b\t", KM)
	fmt.Println("%d\n", MB)
}
