package main

import "fmt"

func main() {
	var name string
	name = "Rafi Septian Hadi"
	fmt.Println(name)

	var firstName = "Rafi"
	fmt.Println(firstName)

	lastName := "Septian"
	fmt.Println(lastName)
	lastName = "Hadi"
	fmt.Println(lastName)

	var age int8 = 20
	fmt.Println(age)

	// auto using int32/64 based on OS
	var realAge = 19
	fmt.Println(realAge)

	// multiple declare variable
	var (
		gender string
		address = "Blablabla"
	)
	gender = "Male"
	fmt.Println(gender)
	fmt.Println(address)
}
