// 1.7 Area of ​​a rhombus with known diagonals

package main

import (
	"fmt"
)

func main() {
	var d1 int
	fmt.Print("> MAJOR DIAGONAL: ")
	fmt.Scan(&d1)

	var d2 int
	fmt.Print("> MINOR DIAGONAL: ")
	fmt.Scan(&d2)

	AreaRombo := d1 * (d2 / 2)

	fmt.Printf(`
	SUM OF THE MAJOR DIAGONAL %d AND MINOR DIAGONAL %d is %d
	`, d1, d2, AreaRombo)
}
