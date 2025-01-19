// 1.8 Volume of a right cylinder.

package main

import (
	"fmt"
)

func main() {
	var r float64
	fmt.Print("> RADIO: ")
	fmt.Scan(&r)

	var h float64
	fmt.Print("> HEIGHT: ")
	fmt.Scan(&h)

	VolCilindro := 3.14159 * (r * r) * h

	fmt.Printf(`
	VOLUMEN OF THE CYLINDER WITH %.2f RADIUS %.2f HEIGHT IS %.2f
	`, r, h, VolCilindro)
}
