package main

import "fmt"

func main() {
	runAppPanic(false)
}

func endApp(){
	// the recover function will catch the panic error
	message := recover()
	if message != nil {
		fmt.Println("Error message:", message)
	}
	fmt.Println("Application was ended")
}

func runAppPanic(error bool) {
	// the defer will called although the app's error
	defer endApp()

	if error{
		panic("App's Error")
		// like throw error, the app will ended here
	}

	fmt.Println("Program's Running")
}
