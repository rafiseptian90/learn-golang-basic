package main

import "fmt"

func AnyType () interface{} {
	return 10
}

func main() {
	var result interface{} = AnyType()

	// check what's type data inside AnyType function
	switch datatype := result.(type){
	case string:
		fmt.Println("Data type is string:", datatype)
	case int:
		fmt.Println("Data type is an integer:", datatype)
	default:
		fmt.Println("Unknown datatype", datatype)
	}
}
