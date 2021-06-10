package main

import "fmt"

func main() {
	// basic switch
	name := "Rafi"
	switch name{
	case "Rafi":
		fmt.Println("Hai ", name)
	case "Septian":
		fmt.Println("Hello ", name)
	default:
		fmt.Println("Unidentified name")
	}

	// short switch
	switch length := len(name); length > 4 {
	case true:
		fmt.Println("True name")
	case false:
		fmt.Println("False name")
	}

	// switch without statement
	score := 80
	switch {
	case score >= 90:
		fmt.Println("Perfect score")
	case score >= 80 && score <= 90:
		fmt.Println("Good score")
	case score <= 80 && score >= 75:
		fmt.Println("Try again")
	default:
		fmt.Println("Go to course")
	}
}
