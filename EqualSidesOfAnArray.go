// Equal Sides Of An Array - 6 kyu
// test#7.1

package main

import (
	"fmt"
	//"strings"
)

func main() {
	arr := []int{1, 100, 50, -51, 1, 1}
	reversIndx := len(arr) - 1

	var firstSum, secondSum int

	for i, num := range arr {

		firstSum += num
		secondSum += arr[reversIndx-i]

		//fmt.Println(firstSum, secondSum)

		if firstSum == secondSum && i == reversIndx-i {
			fmt.Println(i)
			return
		}
		//firstSum += arr[i]
	}

	//fmt.Println(count)
}
