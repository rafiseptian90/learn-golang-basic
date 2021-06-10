package main

import "fmt"

// Filter aliases
type Filter func(string) string

func main() {
	greet("Rafi", filterName)
}

func greet(name string, filter Filter){
	fmt.Println("Hello", filter(name))
}

func filterName(name string) string {
	if name == "blabla" {
		return "..."
	}

	return name
}