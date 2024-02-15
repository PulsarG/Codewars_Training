// Is a number prime? - 6 kyu
// test#8.1

package main

import (
	"fmt"
	//"strings"
)

func main() {
	n := 7

	for i := 2; i < n; i++ {
		if (n % i) == 0 {
			fmt.Println(i, false)
		}
	}

	fmt.Println(true)
}
