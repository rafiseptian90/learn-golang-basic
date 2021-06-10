package main

import "fmt"

func main() {
	initial := getInitial("Rafi", "Septian", "Hadi")

	fmt.Println(initial)
}

func getInitial(firstName string, middleName string, lastName string) string {
	initial := string(firstName[0]) + string(middleName[0]) + string(lastName[0])

	return initial
}
