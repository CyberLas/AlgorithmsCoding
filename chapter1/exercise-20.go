// 1.20 Conversion of degrees Celsius to Fahrenheit, Kelvin and Rankine.

package main

import (
	"fmt"
)

func main() {
	var c float64
	fmt.Print("> GRADE CELSIUS: ")
	fmt.Scan(&c)

	f := (c * 1.8) + 32
	k := c + 273.15
	r := (c * 1.8) + 491.67

	fmt.Printf(`
	CELSIUS TO FARENHEIT, KELVI AND RANKINE:
		%.2f°C => %.2f°F
		%.2f°C => %.2f°K
		%.2f°C => %.2f°R
	`, c, f, c, k, c, r)
}
