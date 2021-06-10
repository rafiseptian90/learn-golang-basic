package main

import "fmt"

func main() {

	months := [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	// take data from index 6 - index 10
	slice1 := months[6:10]
	fmt.Println("Slice 1 = ", slice1)
	fmt.Println("Length of slice1 = ", len(slice1))
	fmt.Println("Capacity of slice1 = ", cap(slice1))

	// append function (add item to last index of slice, if that slice length is full, it will be create a new array)
	appendSlice := append(slice1, "Ramadhan")
	fmt.Println("Append slice = ", appendSlice)

	// make slice (length 2, capacity 5)
	makeSlice := make([]string, 2, 5)
	makeSlice[0] = "Senin"
	makeSlice[1] = "Selasa"
	fmt.Println("Make slice = ", makeSlice)
	
	// copy slice (the length and capacity of slice destination should same	with the source slice from)
	copySlice := make([]string, len(makeSlice), cap(makeSlice))
	copy(copySlice, makeSlice)
	fmt.Println("Copy slice = ", copySlice)

	// this is array
	thisArray := [...]int{1, 2, 3, 4, 5}
	thisSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("This is array = ", thisArray)
	fmt.Println("This is slice = ", thisSlice)

	/*
	Documentation
	- pointer = first index of slice
	- len = we can count total length of that slice
	- cap = total length from start slice index between total length of that array

	!! changing the array or the slice can change the array it self
	*/
}
