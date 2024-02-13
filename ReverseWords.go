// Reverse words 
// (5.2)

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "double spaced words"
	resultString := ""

	words := strings.Fields(str)

	for indx, word := range words {

		runes := []rune(word)

		newString := ""

		for i := len(runes) - 1; i >= 0; i-- {

			newString += string(runes[i])

		} // end inner for

		if indx == len(words)-1 {
			resultString += newString
		} else {
			resultString += newString + " "
		} // end 2 inner for
		
	} // end outer for

	fmt.Println(resultString)
	fmt.Println(resultString == "elbuod decaps sdrow")
}
