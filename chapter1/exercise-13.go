// 1.13 Theorem of cosine

package main

import (
	"fmt"
	"math"
)

func main() {
	var b float64
	fmt.Print("> SIDE 1: ")
	fmt.Scan(&b)

	var c float64
	fmt.Print("> SIDE 2: ")
	fmt.Scan(&c)

	var alfa float64
	fmt.Print("> ANGLE FORMED BY THE KNOWN SIDES: ")
	fmt.Scan(&alfa)

	a := math.Sqrt((b * b) + (c * c) - 2*b*c*math.Cos(alfa*math.Pi/180.0))

	fmt.Printf(`
	THIRD SIDE OF A TRIANGLE HAS A SIDE OF %.2f, %.2f AND AN ANGLE OF %.2f° IS: %.2f
	`, b, c, alfa, a)
}
