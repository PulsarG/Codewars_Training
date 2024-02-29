// The Millionth Fibonacci Kata - 3 kyu
// https://www.codewars.com/kata/53d40c1e2f13e331fc000c26/train/go
// test#18.6

package main

import (
	"fmt"
	"math/big"
	//"strings"
)

// !! Отрицательные
// ** F(n) = F(n+2) - F(n+1)
// ** F(-n) = (-1)**(n+1)*F(n)

// !! Использовать лишь три индекса для экономии памяти

func main() {
	n := 100
	var fibArr [3]big.Int

	if n >= 0 {
		for i := 0; i <= int(n); i++ {
			if i == 0 {
				fibArr[0] = *big.NewInt(int64(i))
				fibArr[1] = *big.NewInt(int64(i))
				fibArr[2] = *big.NewInt(int64(i))
			} else if i == 1 {
				fibArr[1] = *big.NewInt(int64(i))
				fibArr[2] = *big.NewInt(int64(i))
			} else {
				a := fibArr[0]
				b := fibArr[1]

				fibArr[2] = *new(big.Int).Add(&a, &b)
				fibArr[0] = fibArr[1]
				fibArr[1] = fibArr[2]
			}
		}
	} else {
		for i := 0; i >= int(n); i-- {
			if i == 0 {
				fibArr[0] = *big.NewInt(int64(i))
				fibArr[1] = *big.NewInt(int64(i))
				fibArr[2] = *big.NewInt(int64(i))
			} else if i == -1 {
				fibArr[1] = *big.NewInt(int64(i))
				fibArr[2] = *big.NewInt(int64(i))
			} else {
				a := fibArr[0]
				b := fibArr[1]

				fibArr[2] = *new(big.Int).Sub(&a, &b)
				fibArr[0] = fibArr[1]
				fibArr[1] = fibArr[2]
			}
		}
	}
	fmt.Println(fibArr[2])
}
