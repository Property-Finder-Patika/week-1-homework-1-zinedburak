package main

import (
	"fmt"
	"strings"
)

func main() {
	greetPrinter(createGreetInEnglish, "Burak Deniz")
	greetPrinter(createGreetInTurkish, "Burak Deniz")
	greetPrinter(upperCase, "Hello World")

	// Anonymous functions
	func(name string) {
		greeting := "Hello from Anonymous function " + name + " :)"
		fmt.Println(greeting)
	}("Burak Deniz")

	// Giving a name to anonymous function this method is called closure
	closure := func(name string) {
		greeting := "Hello from closure Function " + name + " :)"
		fmt.Println(greeting)
	}
	anotherGreetPrinter(closure, "Burak Deniz")
}

func anotherGreetPrinter(function func(it string), name string) {
	function(name)
}
func createGreetInTurkish(name string) string {
	return "Selam " + name + " :)"
}

func createGreetInEnglish(name string) string {
	return "Hello " + name + " :)"
}
func upperCase(value string) string {
	return strings.ToUpper(value)
}

func greetPrinter(function func(it string) string, name string) {
	var greeting = function(name)
	fmt.Println(greeting)
}
