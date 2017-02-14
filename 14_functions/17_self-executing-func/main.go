package main

import "fmt"

func main() {
	func() {
		fmt.Println("Anonymous self executing function!!!")
	}()
}
