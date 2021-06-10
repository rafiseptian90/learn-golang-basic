package main

import (
	"fmt"
	"learn-golang/beginer/helpers"
)

func main() {
	person := helpers.Person{
		Name:    "Rafi Septian Hadi",
		Age:     19,
		Gender:  "Male",
		Address: "Subang, Jawa Barat",
	}
	fmt.Println(person)
}
