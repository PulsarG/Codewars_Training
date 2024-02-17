// Multiplication table - 6 kyu
// test#11.3+ready

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


// !! BEST ANSWER 

/* func MultiplicationTable(size int) [][]int {
	res := make([][]int, size)
	for i := 0 ; i < size ; i ++ {
	  for x := 1 ; x < size + 1 ; x ++ {
		res[i] = append(res[i], (i + 1) * x)
		}}
	return res
		
  } */
