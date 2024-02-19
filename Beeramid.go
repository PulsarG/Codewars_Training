// Beeramid - 5 kyu
// test#14.3+

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

	leftBeer := floatStartNumber / price
	//leftBeer := int(floatLeftBeer)

	for i := 1; ; i++ {
		a := i
		a *= a
		b := float64(a)
		if b <= leftBeer {
			leftBeer -= b
			count++
		} else {
			fmt.Println(count)
			return
		}
	}
}
