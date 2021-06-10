package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address1 city will not change because address2 isn't reference to address1
	//address2 := address1
	//address2.City = "Jakarta"
	//fmt.Println(address1)
	//fmt.Println(address2)

	// address1 city will change because the address2 reference to address1 using "&"
	address2 := &address1
	address2.City = "Jakarta"
	//fmt.Println(address1)
	//fmt.Println(address2)

	// because the address2 is referenced to address1, so the all data inside address1 will changing too using "*"
	*address2 = Address{"Bandung", "Jawa Barat", "Indonesia"}
	address3 := &address1

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	// using new to copy empty struct
	address4 := new(Address)
	address4.City = "Sukabumi"
	fmt.Println(address4)

	// pointer in function
	changeCountry(&address1, "Germany")
	fmt.Println(address1)
}

func changeCountry(address *Address, country string){
	address.Country = country
}
