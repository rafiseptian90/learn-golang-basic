package helpers

import "fmt"

type Person struct {
	Name string
	Age int
	Gender string
	Address string
}
func (person Person) isMarried() bool {
	return false
}

func sayHello(){
	fmt.Println("Hello")
}