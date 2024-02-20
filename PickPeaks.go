// Pick peaks - 5 kyu
// test#15.4+

package main

import (
	"fmt"
)

var startArr = []int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1}

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

//	PosPeaks{Pos: []int{3, 7, 10}, Peaks: []int{6, 3, 2}}

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
			for i += 2; i < lenArr; i++ {
				if num == startArr[i] {
					continue
				} else if num > startArr[i] {
					isPlatoPeak = true
					break
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
