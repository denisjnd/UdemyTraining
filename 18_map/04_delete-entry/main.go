package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		1 : "Good morning",
		2 : "Good Afternoon",
		3 : "Good Evening",
	}

	fmt.Println(myGreeting)

	// delete an entry
	delete(myGreeting, 3)

	fmt.Println("After deleting: ",myGreeting)

}
