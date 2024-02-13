// Beginner Series #3 Sum of Numbers

package main

import (
	"fmt"
)

func main() {
	var a, b int
	var x, y int

	a = -50
	b = -10

	if a > b {
		x = b
		y = a
	} else {
		x = a
		y = b
	}

	result := x

	for x != y {
		x++
		result += x
	}

	fmt.Println(result)
}
