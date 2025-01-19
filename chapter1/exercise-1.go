// 1.1 Exchange of values.

package main

import "fmt"

func main() {
	var a int
	fmt.Print("> VALUE [A]: ")
	fmt.Scan(&a)

	var b int
	fmt.Print("> VALUE [B]: ")
	fmt.Scan(&b)

	c := a
	a = b
	b = c

	fmt.Printf(`VALUES 
	[A] = %d
	[B] = %d
	`, a, b)
}
