package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyRange       = 30.0
	xyScale       = width / 2 / xyRange
	zScale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30), cos(30)

func main() {
	htmlSVG, err := os.Create("svg.html")
	if err != nil {
		panic(err)
	}
	startString := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	htmlSVG.WriteString(startString)
	fmt.Printf(startString)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			polygonPoint := fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
			htmlSVG.WriteString(polygonPoint)
			fmt.Printf(polygonPoint)
		}
	}
	fmt.Println("</svg>")
	htmlSVG.WriteString("</svg>")
	htmlSVG.Sync()
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyRange * (float64(i)/cells - 0.5)
	y := xyRange * (float64(j)/cells - 0.5)
	// Compute surface height z.

	// Original Shape
	//z := f(x, y)

	// Exercise 3.2
	z := f2(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyScale
	sy := height/2 + (x+y)*sin30*xyScale - z*zScale
	return sx, sy
}
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)

	// Exercise 3.1
	if math.IsNaN(math.Sin(r)/r) || math.IsInf(math.Sin(r)/r, 0) {
		return 0.0
	}
	return math.Sin(r) / r
}

// Exercise 3.2
func f2(x, y float64) float64 {
	return math.Pow(2, math.Sin(x)) * math.Pow(2, math.Sin(y)) / 12
}
