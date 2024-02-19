// Beeramid - 5 kyu
// test#14.2

package main

import (
	"fmt"
	//"strconv"
	//"strings"
	// "unicode"
)

var price float64 = 1.5
var startNumber int = 21

func main() {
	count := 0

	var floatStartNumber = float64(startNumber)

	floatLeftBeer := floatStartNumber / price
	leftBeer := int(floatLeftBeer)

	for i := 1; ; i++ {
		a := i
		a *= a
		if a < leftBeer {
			leftBeer -= a
			count++
		} else {
			fmt.Println(count)
			return
		}
	}
}
