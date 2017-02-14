package main

import "fmt"

func main() {
	name := "Denis John"
	fmt.Println(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>403 Forbidden</title>
		</head>
		<body>

		<h1>` + name + `</h1>

		</body>
		</html>
	`)
}
