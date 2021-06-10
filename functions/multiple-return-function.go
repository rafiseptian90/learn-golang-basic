package main

import "fmt"

func main() {
	name, detail := getPerson()
	fmt.Println("name is", name)
	for i,detail := range detail{
		fmt.Println(i, detail)
	}
}

func getPerson()(string, map[string]string){
	name := "Rafi"
	detail := map[string]string{
		"age" : "19",
		"address": "Subang",
	}

	return name, detail
}
