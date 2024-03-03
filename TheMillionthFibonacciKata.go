// The Millionth Fibonacci Kata - 3 kyu
// https://www.codewars.com/kata/53d40c1e2f13e331fc000c26/train/go
// test#18.9

/* package main

import (
	"fmt"
	"math/big"
	"time"
)

// TODO Сократить, ускорить <12000vc

func main() {
	start := time.Now()
	n := 5
	var fibArr [3]big.Int

	if n >= 0 {
		for i := 0; i <= n; i++ {
			if i == 0 {
				fibArr[0] = *big.NewInt(int64(i))
				fibArr[1] = *big.NewInt(int64(i))
			} else if i == 1 {
				fibArr[1] = *big.NewInt(int64(i))
			} else {
				fibArr[2] = *new(big.Int).Add(&fibArr[0], &fibArr[1])
				fibArr[0] = fibArr[1]
				fibArr[1] = fibArr[2]
			}
		}
	} else {
		for i := -1; i >= n; i-- {
			if i == -1 {
				fibArr[1] = *big.NewInt(int64(i))
			} else {
				fibArr[2] = *new(big.Int).Sub(&fibArr[0], &fibArr[1])
				fibArr[0] = fibArr[1]
				fibArr[1] = fibArr[2]
			}
		}
	}
	fmt.Println(fibArr[2])

	elapsed := time.Since(start)
	fmt.Println("Время выполнения:", elapsed)
} */

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
		for i := 0; i <= n; i++ {
			if i == 0 {
				continue
			} else if i == 1 {
				b = big.NewInt(1)
			} else {
				fib.Add(a, b)
				a.Set(b)
				b.Set(&fib)
			}
		}
	} else {
		for i := 0; i >= n-1; i-- {
			if i == 0 {
				continue
			} else if i == -1 {
				b = big.NewInt(-1)
			} else {
				fib.Sub(b, a)
				b.Set(a)
				a.Set(&fib)
			}
		}
	}

	//fmt.Println(fib)

	elapsed := time.Since(start)
	fmt.Println("Время выполнения:", elapsed)
}
