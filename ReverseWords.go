// Reverse words

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "123 456"
	resultString := ""

	words := strings.Fields(str)

	for _, word := range words {

		runes := []rune(word)

		newString := ""

		for i := len(runes) - 1; i >= 0; i-- {

			newString += string(runes[i])

		} // end inner for

		resultString += newString + " "
	} // end outer for

	fmt.Println(resultString)
}
