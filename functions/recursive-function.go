package main

import "fmt"

func main() {
	result := loopRecursive(5)
	resultRecursive := factorialRecursive(5)

	fmt.Println(result)
	fmt.Println(resultRecursive)
}

// factorial without recursive
func loopRecursive(number int) int {
	result := 1

	for i := number; i > 0 ; i-- {
		result *= i
	}

	return result
}

// factorial with recursive
func factorialRecursive(number int) int {
	if number == 1{
		return 1
	} else {
		return number * factorialRecursive(number - 1)
	}
}
