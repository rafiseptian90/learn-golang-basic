package main

import "fmt"

func main() {
	number := Anything(1)
	strings := Anything("Hi dude")
	boolean := Anything(true)

	fmt.Println(number)
	fmt.Println(strings)
	fmt.Println(boolean)
}

// by using an empty interface, we can use any type data
func Anything(anything interface{}) interface{} {
	return anything
}
