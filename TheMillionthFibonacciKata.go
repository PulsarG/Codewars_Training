// The Millionth Fibonacci Kata - 3 kyu
// https://www.codewars.com/kata/53d40c1e2f13e331fc000c26/train/go
// test#18.8.1

package main

import (
	"fmt"
	"math/big"
)

// TODO Сократить, ускорить <12000vc

func main() {
	n := -9
	var fibArr [2]big.Int

	if n >= 0 {
		for i := 0; i <= int(n); i++ {
			if i == 0 {
				fibArr[0] = *big.NewInt(int64(i))
				fibArr[1] = *big.NewInt(int64(i))
			} else if i == 1 {
				fibArr[1] = *big.NewInt(int64(i))
			} else {
				a := *new(big.Int).Add(&fibArr[0], &fibArr[1])
				fibArr[0] = fibArr[1]
				fibArr[1] = a
			}
		}
	} else {
		for i := -1; i >= int(n); i-- {
			if i == -1 {
				fibArr[1] = *big.NewInt(int64(i))
			} else {
				a := *new(big.Int).Sub(&fibArr[0], &fibArr[1])
				fibArr[0] = fibArr[1]
				fibArr[1] = a
			}
		}
	}
	fmt.Println(fibArr[1])
}
