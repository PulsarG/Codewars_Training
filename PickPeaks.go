// Pick peaks - 5 kyu
// test#15.2

package main

import (
	"fmt"
)

var startArr = []int{1, 2, 3, 6, 4, 1, 2, 3, 2, 1}

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

//	PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}}, */

func main() {
	p := &PosPeaks{}
	
	fmt.Println(searchPicks(startArr, p))
	// ** Поиск плато
	// ** Запись пиков и плато
}

func searchPicks(startArr []int, p *PosPeaks) PosPeaks {
	for i, num := range startArr {
		if i == 0 || i == (len(startArr)-1) {
			continue
		}

		if num > startArr[i-1] && num > startArr[i+1] {
			p.Peaks = append(p.Peaks, num)
			p.Pos = append(p.Pos, i)
		}
	}

	return *p
}

// !! BEST ANSWER
