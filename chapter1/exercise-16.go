// 1.16 Amount to be returned for a borrowed capital

package main

import (
	"fmt"
	"math"
)

func main() {
	var c float64
	fmt.Print("> CAPITAL: ")
	fmt.Scan(&c)

	var t float64
	fmt.Print("> RATE: ")
	fmt.Scan(&t)

	var n float64
	fmt.Print("> NUMBER OF PERIODS: ")
	fmt.Scan(&n)

	m := c * math.Pow((1+(t/100)), n)

	fmt.Printf(`
	LOAN %.2f WITH A RATE %.2f OF WILL BE RETURNED IN A PERIOD %.2f OF %.2f)
	`, c, t, n, m)
}
