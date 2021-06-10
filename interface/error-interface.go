package main

import (
	"bytes"
	"errors"
	"fmt"
)

func main() {
	result, err := Aritmethic(10, "|", 10)

	if err != nil{
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}

func Aritmethic(number1 int, operator string, number2 int) (int, error) {
	if operator == "+" {
		return number1 + number2, nil
	} else if operator == "-" {
		return number1 - number2, nil
	} else if operator == "*" {
		return number1 * number2, nil
	} else if operator == "/" {
		return number1 / number2, nil
	} else {
		message := bytes.Buffer{}
		message.WriteString("Undefined operator: ")
		message.WriteString(operator)

		return 0, errors.New(message.String())
	}
}
