package main

import "fmt"

type Celsius float64

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

type Fahrenheit float64

func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

type Kelvin float64

func (k Kelvin) String() string { return fmt.Sprintf("%gK", k) }

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100

	// Additional Code

	AbsoluteZeroK Kelvin = 0
	FreezingK     Kelvin = 273.15
	BoilingK      Kelvin = 373.15
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func CToK(c Celsius) Kelvin     { return Kelvin(c + 273.15) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func FToK(f Fahrenheit) Kelvin  { return Kelvin(f-32)*5/9 + 273.15 }

func KToC(k Kelvin) Celsius    { return Celsius(k - 273.15) }
func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k-273.15)*9/5 + 32) }

func main() {
	// The goal of the exercise
	// Add types, constants, and functions to tempconv for processing temperatures in the Kelvin scale, where zero Kelvin is -273.15 °C
	// and a difference of 1K has the same magnitude as 1 °C

	// Base code

	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	fmt.Println(AbsoluteZeroC)
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

	// Additional Code
	var k1 Kelvin = -273.15
	var c1 Celsius = 100
	var f1 Fahrenheit = 100

	fmt.Println("Kelvin Constants")
	fmt.Println("Absolute Zero temperature at kevin ", AbsoluteZeroK)
	fmt.Println("Boiling temperature at kelvin ", BoilingK)
	fmt.Println("Freezing temperature at kelvin", FreezingK)

	fmt.Println("Conversion Examples")

	fmt.Printf("%s is %s\n", f1, FToC(f1)) // F to C example 37.7
	fmt.Printf("%s is %s\n", f1, FToK(f1)) // F to K example 310.9

	fmt.Printf("%s is %s\n", c1, CToK(c1)) // C to K example 373.15
	fmt.Printf("%s is %s\n", c1, CToF(c1)) // C to F example 212

	fmt.Printf("%s is %s\n", k1, KToF(k1)) // K to F example == -951.3
	fmt.Printf("%s is %s\n", k1, KToC(k1)) // K to C example == -546.3
}
