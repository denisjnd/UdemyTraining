package main

import (
	"fmt"
	"github.com/denisjnd/UdemyTraining/04_scope/01_package-scope/visibility/vis"
)

func main() {
	fmt.Println("Importing packages")
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
