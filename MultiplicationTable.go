// Multiplication table - 6 kyu
// test#11.2

package main

import (
	"fmt"
	//"strconv"
	//"strings"
	//"unicode"
)

func main() {
	startNumber := 3

	var resultArr [][]int

	for i := 1; i <= startNumber; i++ {

		resultArr = append(resultArr, createNArr(i, startNumber))

	}

	fmt.Println(resultArr)
}

func createNArr(multyplitRate, startNumber int) []int {
	var workArr []int
	for i := 1; i <= startNumber; i++ {
		workArr = append(workArr, i*multyplitRate)
	}
	return workArr
}
