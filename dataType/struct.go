package main

import "fmt"

type Person struct{
	Name, Address, Gender string
	Age int
}

// put a function to a struct
func (person Person) sayHi(name string){
	fmt.Println("Hi", name, "I'm", person.Name)
}

func main() {
	me := Person{
		Name: "Rafi Septian Hadi",
		Gender: "M",
		Address: "Subang",
		Age: 20,
	}
	fmt.Println(me)

	// the param should be structured
	dony := Person{ "Doni Setiawan", "Bandung", "M", 20 }
	fmt.Println(dony)

	// use the sayHi function in Person struct
	renesmee := Person{
		Name: "Renesmee",
		Gender: "F",
		Address: "London",
		Age: 17,
	}
	renesmee.sayHi("Rafi")

}
