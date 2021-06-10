package main

import "fmt"

func main() {
	person := GetPerson("Rafi Septian Hadi")

	if person == nil {
		fmt.Println("Person not detected")
	} else {
		fmt.Println(person)
	}
}

func GetPerson(name string) interface{} {
	if name == "" {
		return nil
	}
	return map[string]string{
		"name": name,
	}
}
