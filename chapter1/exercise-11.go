// 1.11 Area of ​​a triangle as a function of the semiperimeter.

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

	AreaTotalCil := (a + b + c) / 2

	fmt.Printf(`
	AREA OF A TRIANGLE IN A FUNCTION OF THE SEMIPERIMETER OF SIDES %.2f, %.2f, %.2f IS %.2f
	`, a, b, c, AreaTotalCil)
}
