package main

import "C"

//export Sum
func Sum(a int, b int) int {
	return a + b
}

func main() {
	// Need a main function to make CGO compile package as C shared library
}
