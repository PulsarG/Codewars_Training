// Multiples of 3 or 5

package main

import (
	"fmt"
	//"strings"
)

func main() {
	num := 10
	result := 0

	for i := 1; i != num; i++ {
		
		if i%3 == 0 || i%5 == 0 {
			result += i
		}
	}

	fmt.Println(result)
}
