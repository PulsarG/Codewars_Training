package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "1 2 3 0 -3 0"
	intNumbers := [2]int{}

	sings := strings.Fields(s)

	intNumbers[0], _ = strconv.Atoi(sings[0])
	intNumbers[1], _ = strconv.Atoi(sings[0])

	for _, str := range sings {
		if str == " " {
			continue
		} else {

			// to func
			intN, _ := strconv.Atoi(str)

			if intNumbers[0] < intN {
				intNumbers[0] = intN
			}
			if intNumbers[1] > intN {
				intNumbers[1] = intN
			}

		} // end if
	} // end for

	// to func
	fmt.Println(strconv.Itoa(intNumbers[0]) + " " + strconv.Itoa(intNumbers[1]))
}
