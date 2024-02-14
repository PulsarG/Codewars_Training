// Count the divisors of a number
// test#6.2

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


// !! best ANSWER

/* func Divisors(n int)int{
	count:=0
	for i:=1; i <= n; i++{
		if n%i == 0{
			count +=1
		}
	}
	return count
} */