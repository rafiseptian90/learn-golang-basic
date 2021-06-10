package main

import "fmt"

func main() {
	const score int16 = 100
	const score32 = int32(score)

	const name = "Rafi"
	// covert byte from name[0] to string
	fmt.Println(string(name[0]))
}
