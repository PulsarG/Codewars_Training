// Count the divisors of a number
// test#6.1

package main

import (
	"fmt"
	//"strings"
)

func main() {
	n := 500000
	count := 0
	for i := n; i != 0; i-- {
		if n%i == 0 {
			count++
		}
	}
	fmt.Println(count)
}
