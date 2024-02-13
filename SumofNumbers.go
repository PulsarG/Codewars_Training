// Beginner Series #3 Sum of Numbers

package main

import (
	"fmt"
)

func main() {
	var a, b int

	a = -10
	b = -50

	if a > b {
		fmt.Println(GetSum(b, a))
	} else {
		fmt.Println(GetSum(a, b))
	}
}

func GetSum(a, b int) int {
	result := a
	for a != b {
		a++
		result += a
	}
	return result
}
