// 1.12 Hypotenuse of a right triangle.

package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	fmt.Print("> HICK 1: ")
	fmt.Scan(&a)

	var b float64
	fmt.Print("> HICK 2: ")
	fmt.Scan(&b)

	c := math.Sqrt((a * a) + (b * b))

	fmt.Printf(`
	LA HYPOTENUSE OF A TRIANGLE WITH %.2f AND %.2f IS %.2f
	`, a, b, c)
}
