// Build Tower - 6 kyu
// test#12.4+ready

package main

import (
	"fmt"
	//"strconv"
	//"strings"
	// "unicode"
)

func main() {
	startNumber := 10

	for _, str := range TowerBuilder(startNumber) {
		fmt.Println(str)
	}
}

func TowerBuilder(nFloors int) []string {
	var tower []string
	lenStr := ((nFloors * 2) - 1)

	for floor := 1; floor <= nFloors; floor++ {
		var str string
		countStars := ((floor * 2) - 1)
		countSpaces := (lenStr - countStars) / 2

		addSpaces(&str, countSpaces)
		for i := 0; i < countStars; i++ {
			str += "*"
		}
		addSpaces(&str, countSpaces)

		tower = append(tower, str)
	}
	return tower
}

func addSpaces(str *string, countSpaces int) {
	for i := 0; i < countSpaces; i++ {
		*str += " "
	}
}


// !! BEST ANSWER

/* func TowerBuilder(nFloors int) []string {
	tow := make([]string, nFloors)
	for i := 0; i < nFloors; i++ {
	  spaces := strings.Repeat(" ", nFloors - (i + 1))
	  blocks := strings.Repeat("*", i * 2 + 1)
	  tow[i] = spaces + blocks + spaces
	}
	return tow
  } */
