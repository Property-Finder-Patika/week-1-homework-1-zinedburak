package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("Please enter your name:")
	fmt.Scanln(&name)
	greeting := createGreet(name)
	fmt.Printf("%s\n", greeting)
}

func createGreet(name string) string {
	return "Selam " + name + " :)"
}
