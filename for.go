package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan ke-",i)
	}

	fmt.Println("--------------------")

	// loop an slice
	names := []string{"Rafi", "Septian", "Hadi"}
	for i,name := range names{
		fmt.Println("Index-",i," | name:", name)
	}

	fmt.Println("--------------------")

	// loop an slice or array without index
	for _,name := range names{
		fmt.Println(name)
	}

	fmt.Println("--------------------")

	// loop an map
	person := map[string]string{
		"Name" : "Rafi Septian Hadi",
		"Age" : "19",
		"Address": "Subang",
	}
	for key,value := range person{
		fmt.Println(key, "=", value)
	}

}
