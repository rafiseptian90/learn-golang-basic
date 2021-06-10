package main

import (
	"flag"
	"fmt"
)

func main(){
	var host *string = flag.String("host", "localhost", "Put your hostname")
	var user *string = flag.String("user", "root", "Put the user")
	var password *string = flag.String("password", "localhost", "Put the password")

	flag.Parse()

	fmt.Println("Hostname:", *host)
	fmt.Println("User:", *user)
	fmt.Println("Password:", *password)

	// run using: go run flag.go -host="your host" -user="your user" -password="your password"
}