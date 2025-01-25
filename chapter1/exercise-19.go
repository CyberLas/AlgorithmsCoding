// 1.19 Conversion of sexagesimal UºV'W" to hundredths degrees, minutes and seconds

package main

import (
	"fmt"
	"math"
)

func main() {
	var u int
	fmt.Print("> GRADES SEXADECIMAL ANGLE: ")
	fmt.Scan(&u)

	var v int
	fmt.Print("> MINUTES SEXADECIMAL ANGLE: ")
	fmt.Scan(&v)

	var w int
	fmt.Print("> SECONDS SEXADECIMAL ANGLE: ")
	fmt.Scan(&w)

	var s float64
	var c float64

	s = float64(u) + float64(v)/60 + float64(w)/3600
	c = s * (10.0 / 9.0)

	gra := int(c)
	min := int((c - float64(gra)) * 100)
	seg := int(math.Round(((c*100 - float64(gra)*100) - float64(min)) * 100))

	fmt.Printf(`
	ANGLE OF %d° %d' %d" IN SEXAGESIMALS TO HUNDREDTHS IS %d° %d' %d"
	`, u, v, w, gra, min, seg)
}
