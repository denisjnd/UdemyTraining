package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// an artificial input source.
	const input = "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commod consequat. Duis aute irure dolor in reprehenderit in voluptate velit ess cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat no proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	// count the words
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
