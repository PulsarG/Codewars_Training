// Build Tower - 6 kyu
// test#12.1

package main

import (
	"fmt"
	//"strconv"
	//"strings"
	// "unicode"
)

func main() {
	startNumber := 3

	for _, str := range TowerBuilder(startNumber) {
		fmt.Println(str)
	}
}

func TowerBuilder(nFloors int) []string {
	var tower []string
	for i := 0; i < nFloors; i++ {
		var str string

		for j := 0; j < ((i * 2) - 1); j++ {
			str += "*"
		}

		tower = append(tower, str)
	}
	return tower
}
