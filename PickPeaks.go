// Pick peaks - 5 kyu
// https://www.codewars.com/kata/5279f6fe5ab7f447890006a7/train/go
// test#15.5+

package main

import (
	"fmt"
)

var startArr = []int{2, 1, 3, 1, 2, 2, 2, 2}

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

//	 PosPeaks{Pos: []int{2}, Peaks: []int{3}},

func main() {
	p := &PosPeaks{}
	fmt.Println(searchPicks(startArr, p))
}

func searchPicks(startArr []int, p *PosPeaks) PosPeaks {
	lenArr := len(startArr)
	for i := 1; i < lenArr-1; i++ {
		num := startArr[i]
		a := i

		isPeack := (num > startArr[i-1] && num > startArr[i+1])

		// check plato
		isPlatoPeak := false
		if num == startArr[i+1] {
			for i += 1; num <= startArr[i]; i++ {
				if i == lenArr-1 {
					return *p
				}
				if num == startArr[i] {
					isPlatoPeak = true
					continue
				} else if num > startArr[i] {
					isPlatoPeak = true
				}
			} // end for
		} // end if

		if isPeack || isPlatoPeak {
			p.Peaks = append(p.Peaks, num)
			p.Pos = append(p.Pos, a)
		} // end if

	} // end for

	return *p
}

// !! BEST ANSWER
