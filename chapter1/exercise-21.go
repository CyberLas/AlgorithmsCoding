// 1.21 Conversion of ºF, ºK and ºR to Celsius.

package main

import (
	"fmt"
)

func main() {
	var f float64
	fmt.Print("> TEMPERATURE GRADES FARENHEIT: ")
	fmt.Scan(&f)

	var k float64
	fmt.Print("> TEMPERATURE GRADES KELVI: ")
	fmt.Scan(&k)

	var r float64
	fmt.Print("> TEMPERATURE GRADES RANKINE: ")
	fmt.Scan(&r)

	cf := (f - 32) / 1.8
	ck := k - 273.15
	cr := (r - 491.67) / 1.79999999

	fmt.Printf(`
	FARENHEIT, KELVI AND RANKINE TO CELSIUS :
		%.2f°F => %.2f°C
		%.2f°K => %.2f°C
		%.2f°R => %.2f°C
	`, f, cf, k, ck, r, cr)
}
