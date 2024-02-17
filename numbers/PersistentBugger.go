// Persistent Bugger - 6 kyu
// test#10.2+ready

package main

import (
	"fmt"
	"strconv"
	//"strings"
	//"unicode"
)

func fooForTest(num int) int {
	count := 0

	for num >= 10 {
		num = NumToStringAndToTwoNumSum(num)
		count++
	}
	return count
}

func main() {
	num := 25
	fmt.Println(fooForTest(num))
}

func NumToStringAndToTwoNumSum(num int) int {
	str := strconv.Itoa(num)
	result := 1
	for _, char := range str {
		var a int
		a, _ = strconv.Atoi(string(char))
		result *= a
	}
	return result
}

//PersistentBugger
// !! BEST ANSWER

/* func Persistence(n int) int {
	steps := 0
	  for n >= 10 {
	  m := 1
	  for n > 0 {
		m *= n % 10
		n /= 10
	  }
	  n = m
	  steps++
	}
	return steps
  } */
