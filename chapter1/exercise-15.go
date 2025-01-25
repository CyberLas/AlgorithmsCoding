// 1.15 Convert complex c = a + b to polar coordinates

package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	fmt.Print("> REAL PART OF THE COMPLEX: ")
	fmt.Scan(&a)

	var b float64
	fmt.Print("> IMAGINARY PART OF THE COMPLEX: ")
	fmt.Scan(&b)

	w := (math.Pi/2 - math.Atan(a/b)) * 180 / math.Pi
	p := math.Sqrt((a * a) + (b * b))

	fmt.Printf(`
	COMPLEX OF %.2f AND %.2f HAS AS ANGLE %.2f° AND DISTANCE %.2f)
	`, a, b, w, p)
}
