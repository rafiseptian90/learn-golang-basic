package main

import "fmt"

// create HasName interface
type HasName interface{
	GetName() string
}

type Person struct{
	Name, Address string
	Age int
}

type Animal struct{
	Name string
}

// implements the HasName interface
func (person Person) GetName() string{
	return person.Name
}

func (animal Animal) GetName() string{
	return animal.Name
}

func sayHello(object HasName){
	fmt.Println("Hello", object.GetName())
}

func main() {
	person := Person{
		Name: "Rafi Septian Hadi",
		Address: "Subang",
		Age: 19,
	}

	animal := Animal{"Kucing"}

	sayHello(person)
	sayHello(animal)
}
