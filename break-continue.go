package main

import "fmt"

func main() {
	// break
	for i := 1; i <= 10; i++ {
		if i == 6 {
			break
		}

		fmt.Println(i)
	}

	fmt.Println("---------------")

	// continues, will ignore code between this word, but still execute the for loop
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			continue
		}
		// print the odd value
		fmt.Println(i)
	}
}
