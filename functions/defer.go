package main

import "fmt"

func main() {
	result := runApp

	fmt.Println(result(0))
}

func logging(){
	fmt.Println("App was ended")
}

func runApp(number int) int {
	// the defer will called although the app's error
	defer logging()
	result := 10 / number
	return result
}
