package main

import (
	"fmt"
)

func main() {
	i := 4
	print(i)
	var name string = "Betül"
	var greeting = createGreet(name)
	fmt.Printf("%s\n", greeting)
}

func createGreet(name string) string {
	greeting := "Selam " + name + " :)"
	return greeting
}
