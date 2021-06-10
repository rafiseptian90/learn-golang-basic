package main

import "fmt"

func main() {
	counter := 0

	// this is closure (a function inside another function)
	increments := func () {
		counter++
	}
	increments()

	fmt.Println(counter)
}

// we can access variable upside the closure, but we can't access to variable inside closure