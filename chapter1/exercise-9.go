// 1.9 Total area of ​​a right cylinder

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

	AreaTotalCil := 2 * 3.14159 * r * (r + h)

	fmt.Printf(`
	VOLUMEN OF THE RIGHT CYLINDER WITH %.2f RADIUS %.2f HEIGHT IS %.2f
	`, r, h, AreaTotalCil)
}
