// 1.2 Handshake numbers.

package main

import "fmt"

func main() {
	var n int
	fmt.Print("> NUMBER HANDSHAKES: ")
	fmt.Scan(&n)

	a := n * (n - 1) / 2

	fmt.Printf(`
	HANDSHAKES WERE: %d
	`, a)
}
