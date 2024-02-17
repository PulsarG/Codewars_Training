// Build Tower - 6 kyu
// test#12.2+

package main

import (
	"fmt"
	//"strconv"
	//"strings"
	// "unicode"
)

func main() {
	startNumber := 7

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
