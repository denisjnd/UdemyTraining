package main

import "fmt"

func main() {
	records := make([][]string, 0) // make() helps define the length and capacity o a slice
	// student 1
	student1 := make([]string, 4)
	student1[0] = "Denis"
	student1[1] = "John"
	student1[2] = "100.00"
	student1[3] = "74.00"
	// store the record
	records = append(records, student1)
	// student 2
	student2 := make([]string, 4)
	student2[0] = "Denison"
	student2[1] = "Jacob"
	student2[2] = "160.00"
	student2[3] = "740.00"
	// store the record
	records = append(records, student2)
	// student 3
	student3 := make([]string, 4)
	student3[0] = "Denilson"
	student3[1] = "Josh"
	student3[2] = "160.00"
	student3[3] = "740.00"
	// store the record
	records = append(records, student3)
	// print
	fmt.Println(records)
}
