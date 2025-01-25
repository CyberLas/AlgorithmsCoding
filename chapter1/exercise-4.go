// 1.4 Sum of the squares of the first n natural numbers

package main

import "fmt"

func main() {
	var n int
	fmt.Print("> AMOUNT OF NUMBERS: ")
	fmt.Scan(&n)

	s := n * (n + 1) * (2*n + 1) / 6

	fmt.Printf(`
	SUM OF SQUARES FROM 0 TO %d IS: %d
	`, n, s)
}
