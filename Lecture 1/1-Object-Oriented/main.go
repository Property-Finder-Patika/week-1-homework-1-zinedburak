package main

import "fmt"

type Person struct {
	name string
}

func (p Person) greet() string {
	return "Hello " + p.name + " :)"
}
func main() {
	person := Person{name: "Burak Deniz"}
	greeting := person.greet()
	fmt.Println(greeting)
}
