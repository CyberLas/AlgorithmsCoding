// 1.5 Sum of the cubes of the first n natural numbers.

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("> AMOUNT OF NUMBERS: ")
	fmt.Scan(&n)

	s := (n * (n + 1) / 2) * (n * (n + 1) / 2)

	fmt.Printf(`
	SUM OF CUBES FROM 0 TO %d IS: %d
	`, n, s)
}
