// The Millionth Fibonacci Kata - 3 kyu
// https://www.codewars.com/kata/53d40c1e2f13e331fc000c26/train/go
// test#18.2

// Reverse words
// test#5.6

package main

import (
	"fmt"
	"math/big"
	//"strings"
)

// !! Использовать лишь три индекса для экономии памяти

func main() {
	var result big.Int
	n := 2000000
	var fibArr []big.Int

	for i := 0; i <= n; i++ {
		if i == 0 || i == 1 {
			fibArr = append(fibArr, *big.NewInt(int64(i)))
		} else {

			a := fibArr[i-2]
			b := fibArr[i-1]
			c := new(big.Int).Add(&a, &b)
			fibArr = append(fibArr, *c)
		}
	}
	result = fibArr[len(fibArr)-1]

	fmt.Println(result)
}
