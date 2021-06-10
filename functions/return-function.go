package main

import "fmt"

func main() {
	isOdd := isOdd(77)

	if isOdd{
		fmt.Println("This number's odd")
	} else {
		fmt.Println("This number's even")
	}
}

func isOdd(number int) bool {
	if number % 2 == 0 {
		return false
	}

	return true
}
