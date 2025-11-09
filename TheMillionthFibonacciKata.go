// The Millionth Fibonacci Kata - 3 kyu
// https://www.codewars.com/kata/53d40c1e2f13e331fc000c26/train/go
// test#18.17 2

package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	start := time.Now()
	n := 2000000

	if n >= 0 {
		fibAdd := fibAdd(n)
		fmt.Println(fibAdd)
		elapsed := time.Since(start)
		fmt.Println(elapsed)
		return
	} else {
		fibSub := fibSub(n)
		fmt.Println(fibSub)
		elapsed := time.Since(start)
		fmt.Println(elapsed)
	}
}

func fibSub(n int) *big.Int {
	if n == -1 {
		return big.NewInt(-1)
	}

	a, b := big.NewInt(0), big.NewInt(-1)
	for i := -2; i >= n; i-- {
		a.Sub(a, b)
		a, b = b, a
	}
	return b
}

func fibAdd(n int) *big.Int {
	if n <= 0 {
		return big.NewInt(0)
	}

	a, b := big.NewInt(0), big.NewInt(1)
	for i := 2; i <= n; i++ {
		a.Add(a, b)
		a, b = b, a
	}
	return b
}


// Time: 2462ms Passed: 1Failed: 0
Test Results:
Test Example
Basic Tests
Test Passed
Completed in 0.0320ms
You have passed all of the tests! :) 

