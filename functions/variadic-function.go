package main

import "fmt"

func main() {
	fmt.Println(reduce(10, 10, 10))

	// if we want slice as argument
	scores := []int{
		80, 85, 90,
	}
	fmt.Println(reduce(scores...))
}

// variadic must be an final argument
// the argument will converted to slice data type
func reduce(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
