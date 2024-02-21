// Snail - 4 kyu
// https://www.codewars.com/kata/521c2db8ddc89b9b7a0000c1/train/go
// test#16.1

package main

import (
	"fmt"
)

func main() {
	startMap := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// res = []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	result := []int{}

	result = append(result, startMap[0][0])
	result = append(result, startMap[0][1])
	result = append(result, startMap[0][2])

	result = append(result, startMap[1][2])

	result = append(result, startMap[2][2])
	result = append(result, startMap[2][1])
	result = append(result, startMap[2][0])

	result = append(result, startMap[1][0])
	result = append(result, startMap[1][1])

	// 0- 0 1 2
	// 1- 2
	// 2- 2 1 0
	// 1- 0 1

	fmt.Println(result)
}

func Snail(snaipMap [][]int) []int {
	return []int{}
}
