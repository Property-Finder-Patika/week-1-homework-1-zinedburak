package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit  { return Fahrenheit(c*9/5 + 32) }
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func main() {
	// The goal of the exercise
	// Add types, constants, and functions to tempconv for processing temperatures in the Kelvin scale, where zer Kelvin is -273.15 °C
	// and a difference of 1K has the same magnitude as 1 °C

	// Base code

	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F

	// Comparison operators can also be used to compare a value of a named type to another of the same type.
	// or to a value of the underlying type. But two values of different named types cannot be compared directly

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)          // "true"
	fmt.Println(f >= 0)          // "true"
	fmt.Println(c == Celsius(f)) // "true"!

	// Note the last case carefully. In spite of its name, the type conversion Celsius(f) does not change the value of its
	// argument, just its type. The test is true because c and f are both zero

	// You can declare type methods just like they are objects
	// EXAMPLE:

	fmt.Println(c.String()) // "100°C"
	fmt.Printf("%v\n", c)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c)   // "100°C"
	fmt.Println(c)          // "100°C"
	fmt.Printf("%g\n", c)   // "100"; does not call String
	fmt.Println(float64(c)) // "100"; does not call String

}
