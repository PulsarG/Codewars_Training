// Pick peaks - 5 kyu
// https://www.codewars.com/kata/5279f6fe5ab7f447890006a7/train/go
// test#15.6+ready

package main

import (
	"fmt"
)

var startArr = []int{5, 7, 6, 6, -1, -1, -1, 10, 10, 10, 10, 10}

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

//	 <kata.PosPeaks>: {Pos: [1, 5], Peaks: [7, 10]}

func main() {
	p := &PosPeaks{}
	fmt.Println(searchPicks(startArr, p))
}

func searchPicks(startArr []int, p *PosPeaks) PosPeaks {
	lastIndex := len(startArr) - 1

	for i := 1; i < lastIndex; i++ {
		num := startArr[i]
		currentIndexOrStartPlato := i

		isPeack := (num > startArr[i-1] && num > startArr[i+1])

		// check plato
		isPlatoPeak := false
		if num == startArr[i+1] && num > startArr[i-1] {
			for i += 1; num <= startArr[i]; i++ {
				if i == lastIndex {
					return *p
				} else if num == startArr[i] {
					isPlatoPeak = true
					continue
				} else if num < startArr[i] {
					isPlatoPeak = false
					i -= 1
					break
				} else if num > startArr[i] {
					isPlatoPeak = true
					break
				}
			} // end for
		} // end if

		if isPeack || isPlatoPeak {
			p.Peaks = append(p.Peaks, num)
			p.Pos = append(p.Pos, currentIndexOrStartPlato)
		} // end if

	} // end for

	return *p
}

// !! BEST ANSWER

/* func PickPeaks(array []int) PosPeaks {
	posPeaks := PosPeaks{Pos: []int{}, Peaks: []int{}}
	candidate := 0
	for i := 1; i < len(array); i++ {
		if array[i-1] < array[i] {
			candidate = i
		} else if array[i-1] > array[i] && candidate > 0 {
			posPeaks.Pos = append(posPeaks.Pos, candidate)
			posPeaks.Peaks = append(posPeaks.Peaks, array[candidate])
			candidate = 0
		}
	}
	return posPeaks
} */