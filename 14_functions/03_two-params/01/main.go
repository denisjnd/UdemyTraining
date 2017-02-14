package main

import "fmt"

func main() {
	greet("John", "Davis")
}

func greet(fname, lname string) { // since there are of they same type. they dont need coma
	fmt.Println(fname, lname)
}

//func greet(fname string, lname string) {
//	fmt.Println(fname, lname)
//}
