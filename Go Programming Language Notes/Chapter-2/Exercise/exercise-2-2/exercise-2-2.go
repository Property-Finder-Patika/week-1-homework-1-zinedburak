package main

import (
	"fmt"
	"os"
	"strconv"
)

type Meter float64
type Feet float64

func (m Meter) String() string { return fmt.Sprintf("%g M", m) }
func (f Feet) String() string  { return fmt.Sprintf("%g Feet", f) }

func MToF(meter Meter) Feet { return Feet(meter * 3.29) }
func FToM(feet Feet) Meter  { return Meter(feet / 3.29) }

func main() {
	// Write a general-purpose unit conversion program analogous to cf that reads numbers
	// from its command line arguments or from the standard input if there are no arguments,
	// and converts each number into units like temperature in Celsius and Fahrenheit,
	// length in feet and meters, weight in pounds and kilograms
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf : %v\n", err)
			os.Exit(1)
		}
		f := Feet(t)
		m := Meter(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, FToM(f), m, MToF(m))
	}
}
