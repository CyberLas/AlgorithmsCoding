// 1.3 Sum of the first n natural numbers

package main

import "fmt"

func main() {
	var n int
	fmt.Print("> AMOUNT OF NUMBERS: ")
	fmt.Scan(&n)

	s := n * (n + 1) / 2

	fmt.Printf(`
	SUM OF NUMBERS OF 0 TO %d IS: %d
	`, n, s)
}
