package main

import "fmt"

func main() {
	book := map[string]string{
		"title" : "The Twilight Saga: Eclipse",
		"author": "Stephenie Meyer",
		"year": "2006",
	}

	fmt.Println(book["title"])
	fmt.Println(book["author"])
	fmt.Println(book["year"])

	// count length data
	fmt.Println("Length book data = ", len(book))

	// delete key
	delete(book, "year")
	fmt.Println(book)

	// manually create map
	var person = make(map[string]string)
	person["name"] = "Rafi Septian Hadi"
	person["age"] = "19"
	fmt.Println(person)
}
