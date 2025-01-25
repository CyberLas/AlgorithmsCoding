// 1.14 Distance between 2 points

package main

import (
	"fmt"
	"math"
)

func main() {
	var x1 float64
	fmt.Print("> ABSCISSA OF POINT 1: ")
	fmt.Scan(&x1)

	var y1 float64
	fmt.Print("> ORDINATE OF POINT 1: ")
	fmt.Scan(&y1)

	var x2 float64
	fmt.Print("> ABSCISSA OF POINT 2: ")
	fmt.Scan(&x2)

	var y2 float64
	fmt.Print("> ORDINATE OF POINT 2: ")
	fmt.Scan(&y2)

	d := math.Sqrt(((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1)))

	fmt.Printf(`
	THE DISTANCE BETWEEN POINT (%.2f, %.2f) AND POINT (%.2f, %.2f) IS: %.2f
	`, x1, y1, x2, y2, d)
}
