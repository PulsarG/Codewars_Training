// Multiples of 3 or 5

package main

import (
	"fmt"
	//"strings"
)

func main() {
	num := 10
	result := 0

	for num != 0 {
		num--
		if num%3 == 0 || num%5 == 0 {
			result += num
		}
	}

	// OR (bad)
	/* for i := 1; i != num; i++ {
		if i%3 == 0 || i%5 == 0 {
			result += i
		}
	} */

	fmt.Println(result)
}


// !! BETTER ANSWER

/* func Multiple3And5(n int) int {
	sum := 0;
	for i :=1; i < n; i++ {
		if (i % 3 == 0) || (i % 5 == 0) {
			sum += i
		}
	}
	return sum;
} */