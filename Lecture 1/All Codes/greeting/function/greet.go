package main

import (
	"fmt"
)

func main() {
	greet("Zeynep")
}

func greet(name string) {
	fmt.Printf("Selam %s :)\n", name)
}
