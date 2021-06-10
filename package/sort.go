package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age int
}

// Users create aliases for array of users
type Users []User

func (users Users) Len() int {
	return len(users)
}

func (users Users) Less(i, j int) bool {
	return users[i].Age < users[j].Age
}

func (users Users) Swap(i, j int) {
	users[i], users[j] = users[j], users[i]
}

func main(){
	users := []User{
		{ "Rafi", 19 },
		{ "Ahmad", 21 },
		{ "Agung", 20 },
		{ "Hana", 22 },
	}

	// before sort
	fmt.Println(users)
	// because the users not have contract with sort interface, so we should parse using Users type
	sort.Sort(Users(users))
	// after sort
	fmt.Println(users)
	// length of users
	fmt.Println(Users.Len(users))
}