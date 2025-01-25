// 1.6 Sum of the areas of all the squares formed by joining the midpoints of their sides

package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Print("> NUMBER OF MIDPOINTS: ")
	fmt.Scan(&a)

	s := 2 * (a * a)

	fmt.Printf(`
	SUM OF THE AREAS OF A SQUARE BY JOINING ITS %d IS: %d
	`, a, s)
}
