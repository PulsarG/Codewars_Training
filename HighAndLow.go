package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "1 2 3"
	intNumbers := [2]int{0, 0}
	resultS := ""

	sings := strings.Fields(s)

	intNumbers[0], _ = strconv.Atoi(sings[0])
	intNumbers[1], _ = strconv.Atoi(sings[0])

	for _, str := range sings {
		if str == " " {
			continue
		} else {

			intN, _ := strconv.Atoi(str)

			if intNumbers[0] <= intN {
				intNumbers[0] = intN
			}
			if intNumbers[1] >= intN {
				intNumbers[1] = intN
			}

		} // end if
	} // end for

	resultS += strconv.Itoa(intNumbers[0]) + " " + strconv.Itoa(intNumbers[1])

	fmt.Println(resultS)
}
