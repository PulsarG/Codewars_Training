// Beeramid - 5 kyu
// test#14.4+ready

package main

import (
	"fmt"
)

var price float64 = 1.5
var startNumber int = 21

func main() {
	count := 0
	leftBeer := float64(startNumber) / price

	for i := 1; ; i++ {
		b := float64(i * i)

		if b <= leftBeer {
			leftBeer -= b
			count++
		} else {
			fmt.Println(count)
			return
		}
	}
}

// !! BEST ANSWER

/* func Beeramid(bonus int, price float64) (row int) {
	cans := int(float64(bonus) / price)
	for cans > 0 {
		cans = cans - row*row
		if cans >= (row+1)*(row+1) {
			row++
		}
	}
	return row
} */
