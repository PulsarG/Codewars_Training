// Equal Sides Of An Array - 6 kyu
// test#7.1-ready+

package main

import (
	"fmt"
	//"strings"
)

func main() {
	arr := []int{1, 100, 50, -51, 1, 1}
	//reversIndx := len(arr) - 1

	//var firstSum, secondSum int

	for idxForFind := 0; idxForFind != len(arr); idxForFind++ {

		var firstSum, secondSum int

		for idxDown := (idxForFind - 1); idxDown >= 0; idxDown-- {
			firstSum += arr[idxDown]
		} // end for

		for idxUp := (len(arr) - 1); idxUp != idxForFind; idxUp-- {
			secondSum += arr[idxUp]
		} // end for

		if firstSum == secondSum {
			fmt.Println("Result: ", idxForFind)
		} // ned if

	} // end outer for
}


// !! ANSWER

/* func FindEvenIndex(arr []int) int {
	var sum, b int
	for _,n := range arr { sum += n }
	for i,n := range arr {
	  sum -= n
	  if sum == b { return i }
	  b += n
	}
	return -1
  } */
