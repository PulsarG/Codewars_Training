// The Millionth Fibonacci Kata - 3 kyu
// https://www.codewars.com/kata/53d40c1e2f13e331fc000c26/train/go
// test#18.2

// Reverse words
// test#5.6

package main

import (
	"fmt"
	//"math/big"
	//"strings"
)

func main() {
	var result int
	n := 50
	var fibArr []int

	for i := 0; i <= n; i++ {
		if i == 0 || i == 1 {
			fibArr = append(fibArr, i)
		} else {
			fibArr = append(fibArr, (fibArr[i-2] + fibArr[i-1]))
		}
	}
	result = fibArr[len(fibArr)-1]

	fmt.Println(result)
}
