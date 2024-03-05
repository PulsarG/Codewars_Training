// The Millionth Fibonacci Kata - 3 kyu
// https://www.codewars.com/kata/53d40c1e2f13e331fc000c26/train/go
// test#18.11

/* package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	start := time.Now()
	n := 1000000 // 2000000
	var fib big.Int

	a, b := big.NewInt(0), big.NewInt(0)

	if n >= 0 {
		for i := 0; i <= n; i++ {
			if i == 0 {
				fib.Add(a, b)
			} else if i == 1 {
				b = big.NewInt(1)
				fib.Add(a, b)
			} else {
				fib.Add(a, b)
				a.Set(b)
				b.Set(&fib)
			}
		}
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

	//fmt.Println(fib)

	elapsed := time.Since(start)
	fmt.Println("Время выполнения:", elapsed)
} */

package main

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

func main() {
	start := time.Now()

	n := 2000000

	// Бине
	phi := (1 + math.Sqrt(5)) / 2
	fib := (math.Pow(phi, float64(n)) - math.Pow(-phi, -float64(n))) / math.Sqrt(5)

	result := big.NewInt(int64(fib))

	fmt.Println(result)

	elapsed := time.Since(start)
	fmt.Println("Время выполнения:", elapsed)
}
