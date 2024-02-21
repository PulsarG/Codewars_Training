// Snail - 4 kyu
// https://www.codewars.com/kata/521c2db8ddc89b9b7a0000c1/train/go
// test#16.2

package main

import (
	"fmt"
)

func main() {
	startMap := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	res := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	result := []int{}
	isWork := true

	for i := 0; i < len(startMap[0]); i++ {
		result = append(result, startMap[0][i])
	}

	//startMap = append(startMap[:0], startMap[0+1:]...)
	/* result = append(result, startMap[0][0])
	result = append(result, startMap[0][1])
	result = append(result, startMap[0][2])
	*/
	result = append(result, startMap[1][2])

	result = append(result, startMap[2][2])
	result = append(result, startMap[2][1])
	result = append(result, startMap[2][0])

	result = append(result, startMap[1][0])
	result = append(result, startMap[1][1])

	startMap = append(startMap[:0], startMap[0+1:]...)

	for i, num := range result {
		if num != res[i] {
			isWork = false
		}
	}
	fmt.Println(isWork, startMap)
}
