// The Millionth Fibonacci Kata - 3 kyu
// https://www.codewars.com/kata/53d40c1e2f13e331fc000c26/train/go
// test#18.12

package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	start := time.Now()
	n := 2000000
	var fib big.Int

	a, b := big.NewInt(0), big.NewInt(0)

	if n >= 0 {
		fibAdd := fibA(n)
		fmt.Println(fibAdd)
		elapsed := time.Since(start)
		fmt.Println("Время выполнения:", elapsed)
		return
	} else {
		for i := 0; i >= n-1; i-- {
			if i == 0 {
				fib.Sub(b, a)
			} else if i == -1 {
				b = big.NewInt(-1)
				fib.Sub(b, a)
			} else {
				fib.Sub(b, a)
				b.Set(a)
				a.Set(&fib)
			}
		}
	}

	fmt.Println(fib)

	elapsed := time.Since(start)
	fmt.Println("Время выполнения:", elapsed)
}

func fibA(n int) *big.Int {
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
