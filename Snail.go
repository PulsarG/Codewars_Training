// Snail - 4 kyu
// https://www.codewars.com/kata/521c2db8ddc89b9b7a0000c1/train/go
// test#16.8+

package main

import (
	"fmt"
)

func main() {
	startMap := [][]int{{1, 2, 3, 4, 5}, {16, 17, 18, 19, 6}, {15, 24, 25, 20, 7}, {14, 23, 22, 21, 8}, {13, 12, 11, 10, 9}}
	//res := []int{1, 2, 3, 1, 4, 7, 7, 9, 8, 7, 7, 4, 5, 6, 9, 8}
	result := []int{}
	isWork := true

	for len(startMap) >= 1 {
		l := len(startMap)
		// запись первого слайса
		for i := 0; i < len(startMap[0]); i++ {
			result = append(result, startMap[0][i])
		}
		if l > 1 {
			// запись последних элементов всех слайсов далее
			for i := 1; i < l; i++ {
				result = append(result, startMap[i][l-1])
				startMap[i] = startMap[i][:(l - 1)]
			}
			startMap = append(startMap[:0], startMap[0+1:]...)
			l = len(startMap)
			// отзеркаливание
			for i := 0; i < l/2; i++ {
				startMap[i], startMap[l-i-1] = startMap[l-i-1], startMap[i]
			} // end for
			for i := 0; i < l; i++ {
				for j := 0; j < l/2; j++ {
					startMap[i][j], startMap[i][l-j-1] = startMap[i][l-j-1], startMap[i][j]
				}
			} // end for
		} else if l == 1 {
			startMap = append(startMap[:0], startMap[0+1:]...) // end for
		} // end if
	} // end for

	// check
	/* for i, num := range result {
		if num != res[i] {
			isWork = false
		}
	} */
	fmt.Println(isWork, result)
}
