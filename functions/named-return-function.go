package main

import "fmt"

func main() {
	first, middle, last := getFullname()
	fmt.Println(first)
	fmt.Println(middle)
	fmt.Println(last)
}

func getFullname()(firstName string, middleName string, lastName string){
	firstName = "Rafi"
	middleName = "Septian"
	lastName = "Hadi"

	return
}
