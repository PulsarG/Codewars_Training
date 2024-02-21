package main

import (
	"fmt"
)

func main() {
	a := 10000

	for i := 0; i < 2; i++ {
		a += (a / 10)
	}

	fmt.Println((a - 10000))
}
