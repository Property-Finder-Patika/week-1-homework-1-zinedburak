package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	file, _ := os.Create("Newton.png")
	defer file.Close()
	const (
		xMin, yMin, xMax, yMax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(yMax-yMin) + yMin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xMax-xMin) + xMin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newtonMethod(z))
		}
	}

	png.Encode(file, img) // NOTE: ignoring errors
}
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			// Exercise 3.5
			switch {
			case n > 50:
				return color.RGBA{100, 0, 0, 255}
			default:
				return color.RGBA{0, 0, 255, 117}
			}
			// Original
			//return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

// Exercise 3.7
func newtonMethod(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
