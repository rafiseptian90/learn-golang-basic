package main

import (
	"fmt"
	"regexp"
)

func main(){
	// must be initiated with "y" and must be end with "a"
	regex := regexp.MustCompile("y([a-z])([a-z])a")
	fmt.Println(regex.MatchString("yena"))
	fmt.Println(regex.MatchString("yeni"))

	searchStrings := regex.FindAllString("yena yani yuna", -1)
	fmt.Println(searchStrings)
}
