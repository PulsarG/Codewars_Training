package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "1 2 3 4 5 6 -1"
	intNumbers := [2]int{0, 0}
	resultS := ""

	words := strings.Fields(s)

	for _, str := range words {
		if str == " " {
			continue
		} else {

			intN, _ := strconv.Atoi(str)

			if intNumbers[0] < intN {
				intNumbers[0] = intN
			}
			if intNumbers[1] > intN {
				intNumbers[1] = intN
			}

		} // end if
	} // end for

	for _, num := range intNumbers {
		resultS += strconv.Itoa(num) + " "
	}

	fmt.Println(resultS)
}
