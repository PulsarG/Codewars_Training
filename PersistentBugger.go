// Persistent Bugger - 6 kyu
// test#10.1

package main

import (
	"fmt"
	"strconv"
	//"strings"
	//"unicode"
)

func main() {
	num := 25
	count := 0

	for num >= 10 {
		num = StringToTwoNumSum(num)
		count++
	}

	fmt.Println(count)
}

func StringToTwoNumSum(num int) int {
	str := strconv.Itoa(num)
	result := 1
	for _, char := range str {
		var a int
		a, _ = strconv.Atoi(string(char))
		result *= a
	}
	return result
}
