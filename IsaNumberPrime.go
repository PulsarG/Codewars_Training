// Is a number prime? - 6 kyu
// test#8.2+ready

package main

import (
	"fmt"
	"math"
	//"strings"
)

func main() {
	n := 8438644364383
	if n == 2 || n == 3 || n == 5 || n == 7 { fmt.Println(true) }
  
if n <= 1 { fmt.Println(false) }

maxDivisor := int(math.Sqrt(float64(n))) + 1 
  
for i := 2; i <= maxDivisor; i++ {
  if i > 3 && i%2 == 0 {continue}
  if i > 3 && i%3 == 0 {continue}
		if (n % i) == 0 {
			fmt.Println(false)
		}
	}

	fmt.Println(true)
}
