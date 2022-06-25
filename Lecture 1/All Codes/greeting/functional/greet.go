package main

import (
	"fmt"
	"strings"
)

func main() {
	greetPrinter(createGreetInTurkish, "Hatice")
	greetPrinter(createGreetInEnglish, "Mary")
	greetPrinter(convertToUppercase, "naber?")

	greetCreator := createGreetInTurkish
	greetPrinter(greetCreator, "Zeynep")

	func(name string) {
		greeting := "Selam " + name + " :)"
		fmt.Printf("%s\n", greeting)
	}("Yesim")

	closure := func(name string) {
		greeting := "Selam " + name + " :)"
		fmt.Printf("%s\n", greeting)
	}
	closure("Fatma")
	anotherGreetPrinter(closure, "Zeynep")
}

func createGreetInTurkish(name string) string {
	return "Selam " + name + " :)"
}

func createGreetInEnglish(name string) string {
	return "Hi " + name + " :)"
}

func convertToUppercase(arg string) string {
	return strings.ToUpper(arg)
}

func greetPrinter(function func(it string) string, name string) {
	var greeting = function(name)
	fmt.Printf("%s\n", greeting)
}

func anotherGreetPrinter(function func(it string), name string) {
	function(name)
}

/*func add(x int) func(y int) int {
	return func(y int) int {
		return x + y
	}
}*/
