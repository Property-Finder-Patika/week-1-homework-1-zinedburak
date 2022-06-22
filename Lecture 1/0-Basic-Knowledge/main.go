package main

import (
	"fmt"
)

//These are the allowed variable declarations outside any function
//	var yourInt = 200
//	var yourInt int

//You can not use this method out of functions
//	yourInt := 200

func main() {

	fmt.Println("Hello World ")
	var name = "Burak Deniz"
	var greeting = createGreet(name)
	fmt.Println(greeting)
}

func createGreet(name string) string {
	greeting := "Hello " + name + " :)"
	return greeting
}
