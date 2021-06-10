package main

import "fmt"

func main() {
	// manual input array
	var names [3]string
	names[0] = "Rafi"
	names[1] = "Septian"
	names[2] = "Hadi"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// auto fill array
	scores := [3]int{
		90,
		95,
		99,
	}
	fmt.Println(scores)
	fmt.Println(scores[0])
	fmt.Println(scores[1])
	fmt.Println(scores[2])

	// count length of array (not value length of array)
	fmt.Println(len(scores))

	/*
	Documentation
	- The value of array must the same data type
	- if we try to count length of empty array, it will returning the length of that array, not the length of values
	*/
}
