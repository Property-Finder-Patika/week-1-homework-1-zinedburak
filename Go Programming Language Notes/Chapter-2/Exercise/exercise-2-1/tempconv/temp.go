package tempconv

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
