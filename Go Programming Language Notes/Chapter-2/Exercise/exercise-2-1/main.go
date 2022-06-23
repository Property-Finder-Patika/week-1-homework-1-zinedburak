package main

import (
	"fmt"
	"tempconv/exercise-2-1/tempconv"
)

func main() {
	// The goal of the exercise
	// Add types, constants, and functions to tempconv for processing temperatures in the Kelvin scale, where zero Kelvin is -273.15 °C
	// and a difference of 1K has the same magnitude as 1 °C

	// Base code

	fmt.Printf("%g\n", tempconv.BoilingC-tempconv.FreezingC) // "100" °C
	fmt.Println(tempconv.AbsoluteZeroC)
	boilingF := tempconv.CToF(tempconv.BoilingC)
	fmt.Printf("%g\n", boilingF-tempconv.CToF(tempconv.FreezingC)) // "180" °F

	// Comparison operators can also be used to compare a value of a named type to another of the same type.
	// or to a value of the underlying type. But two values of different named types cannot be compared directly

	var c tempconv.Celsius
	var f tempconv.Fahrenheit
	fmt.Println(c == 0)                   // "true"
	fmt.Println(f >= 0)                   // "true"
	fmt.Println(c == tempconv.Celsius(f)) // "true"!

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

	// Additional Code
	var k1 tempconv.Kelvin = -273.15
	var c1 tempconv.Celsius = 100
	var f1 tempconv.Fahrenheit = 100

	fmt.Println("Kelvin Constants")
	fmt.Println("Absolute Zero temperature at kevin ", tempconv.AbsoluteZeroK)
	fmt.Println("Boiling temperature at kelvin ", tempconv.BoilingK)
	fmt.Println("Freezing temperature at kelvin", tempconv.FreezingK)

	fmt.Println("Conversion Examples")

	fmt.Printf("%s is %s\n", f1, tempconv.FToC(f1)) // F to C example 37.7
	fmt.Printf("%s is %s\n", f1, tempconv.FToK(f1)) // F to K example 310.9

	fmt.Printf("%s is %s\n", c1, tempconv.CToK(c1)) // C to K example 373.15
	fmt.Printf("%s is %s\n", c1, tempconv.CToF(c1)) // C to F example 212

	fmt.Printf("%s is %s\n", k1, tempconv.KToF(k1)) // K to F example == -951.3
	fmt.Printf("%s is %s\n", k1, tempconv.KToC(k1)) // K to C example == -546.3
}
