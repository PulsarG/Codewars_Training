// Snail - 4 kyu
// https://www.codewars.com/kata/521c2db8ddc89b9b7a0000c1/train/go
// test#16.4

package main

import (
	"fmt"
)

func main() {
	startMap := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	res := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	result := []int{}
	isWork := true

	for len(startMap) >= 1 {

		// !! НАЧАЛО ЦИКЛА
		l := len(startMap)

		// запись первого слайса
		for i := 0; i < len(startMap[0]); i++ {
			result = append(result, startMap[0][i])
		}

		// запись последних элементов всех слайсов далее
		for i := 1; i < l; i++ {
			result = append(result, startMap[i][l-1])
			startMap[i] = startMap[i][:(l - 1)]
		}

		startMap = append(startMap[:0], startMap[0+1:]...)
		l = len(startMap)

		if l > 0 {
			// отзеркаливание
			for i := 0; i < l/2; i++ {
				for j := 0; i < l/2; i++ {
					startMap[i][j], startMap[i][l-j-1] = startMap[i][l-j-1], startMap[i][j]
				}
				startMap[i], startMap[l-i-1] = startMap[l-i-1], startMap[i]
			} // end for
		} // end if
	} // end for

	// check
	for i, num := range result {
		if num != res[i] {
			isWork = false
		}
	}
	fmt.Println(isWork, startMap, result)
}