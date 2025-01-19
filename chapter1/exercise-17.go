// 1.17 Law of sines.

package main

import (
	"fmt"
	"math"
)

func main() {
	var c float64
	fmt.Print("> HICK: ")
	fmt.Scan(&c)

	var alfa float64
	fmt.Print("> ANGLE 1: ")
	fmt.Scan(&alfa)

	var beta float64
	fmt.Print("> ANGLE 2: ")
	fmt.Scan(&beta)

	var gamma float64
	fmt.Print("> ANGLE 3: ")
	fmt.Scan(&gamma)

	a := (c * math.Sin((beta * math.Pi / 180))) / math.Sin((alfa * math.Pi / 180))
	b := (c * math.Sin((gamma * math.Pi / 180))) / math.Sin((alfa * math.Pi / 180))

	fmt.Printf(`
	SIDE OPPORSITE ANGLE %.2f IS : %.2f
	SIDE OPPORSITE ANGLE %.2f IS : %.2f
	SIDE OPPORSITE ANGLE %.2f IS : %.2f
	`, alfa, c, gamma, a, beta, b)
}
