package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `required:"true" min:"3"`
	Age int `min:"10" max:"70"`
}

func main(){
	person := Person{"Rafi", 19}

	if isValidStruct(person){
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}
}

func isValidStruct(data interface{}) bool {
	// grab length of this struct
	for i := 0; i < reflect.TypeOf(data).NumField() ; i++ {
		// get each field
		field := reflect.TypeOf(data).Field(i)
		// required validation
		if field.Tag.Get("required") == "true" {
			// false if the required field isn't set, true if the required field is was set
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}

	return true
}
