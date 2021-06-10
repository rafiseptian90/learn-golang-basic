package main

import "fmt"

func main() {
	// declare aliases
	type Age int8
	type isMarried bool

	// example
	const age Age = 20
	const statusMarried isMarried = false
	
	fmt.Println(age)
	fmt.Println(statusMarried)
}
