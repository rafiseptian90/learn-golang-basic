package main

import "fmt"

func main() {
	name := "Septian"

	if name == "Rafi" {
		fmt.Println("Hai ", name)
	} else if name == "Septian"{
		fmt.Println("Hello ", name)
	} else {
		fmt.Println("Unidentified name")
	}

	// short statement
	if score := 90; score > 85 {
		fmt.Println("Great score")
	} else if score <= 75 && score >= 70 {
		fmt.Println("Try again")
	} else {
		fmt.Println("Go to course")
	}
}
