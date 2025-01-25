// 1.18 Conversion of sexagesimal degrees to hundredths and radians

package main

import (
	"fmt"
	"math"
)

func main() {
	var s float64
	fmt.Print("> NUMBER IN SEXADECIMAL: ")
	fmt.Scan(&s)

	c := (10 * s) / 9
	r := (math.Pi * s) / 180

	fmt.Printf(`
	SEXAGESIMAL %.2f TO CENTENSIMALS IS %.2f
	SEXAGESIMAL %.2f TO RADIANS IS %.2f
	`, s, c, s, r)
}
