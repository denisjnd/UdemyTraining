package main

import "fmt"

func main() {
	half := func(n int) (int, bool) {
		return n /2, n%2 == 0
	}
	fmt.Println(half(5))
}


//package main
//
//import "fmt"
//
//func half(x int) (int, bool) {
//
//	div := x / 2
//
//	if x%2 == 0 {
//		return div, true
//	} else {
//		return div, false
//	}
//
//}
//
//func main() {
//	fmt.Println(half(3))
//}
