package main

import (
	"fmt"
)

type Blacklist func(string) bool

func main() {
	// wrap anonymous function to a variable
	blacklisted := func(name string) bool {
		blacklistedNames := []string{
			"admin",
			"blabla",
			"root",
		}

		for _, blacklistedName := range blacklistedNames {
			if blacklistedName == name {
				return true
			}
		}

		return false
	}
	registerUser("Rafi", blacklisted)

	// pure anonymous
	registerUser("admin", func(name string) bool {
		blacklistedNames := []string{
			"admin",
			"blabla",
			"root",
		}

		for _, blacklistedName := range blacklistedNames {
			if blacklistedName == name {
				return true
			}
		}

		return false
	})
}

func registerUser(name string, blacklisted Blacklist){
	if blacklisted(name){
		fmt.Println("You has been blacklisted", name)
	} else {
		fmt.Println("Welcome", name)
	}
}
