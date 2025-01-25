// 1.10 Area of ​​a triangle inscribed in a circle

package main

import (
	"fmt"
)

func main() {
	var a float64
	fmt.Print("> SIDE A: ")
	fmt.Scan(&a)

	var b float64
	fmt.Print("> SIDE B: ")
	fmt.Scan(&b)

	var c float64
	fmt.Print("> SIDE C: ")
	fmt.Scan(&c)

	var r float64
	fmt.Print("> RADIUS: ")
	fmt.Scan(&r)

	AreaTotalCil := ((a + b + c) / 2) * r

	fmt.Printf(`
	AREA OF A TRIANGLE INSCRIBED IN A CIRCLE WITH SIDE %.2f, %.2f, %.2f AND RADIUS %.2f IS %.2f
	`, a, b, c, r, AreaTotalCil)
}
