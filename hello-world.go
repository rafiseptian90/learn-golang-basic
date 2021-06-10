package main

import "fmt"

func main(){
	greeting("Rafi", "Septian Hadi")
}

func greeting(firstName string, lastName string)  {
	fmt.Println(firstName + " " + lastName)
}