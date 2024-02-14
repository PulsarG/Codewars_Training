// Reverse words
// test#5.6

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "double"
	resultString := ""

	words := strings.Fields(str)

	countSpaces := strings.Count(str, " ")
	if countSpaces != 0 {
		countSpaces /= (len(words) - 1)
	}

	for indx, word := range words {

		runes := []rune(word)

		newString := ""

		for i := len(runes) - 1; i >= 0; i-- {

			newString += string(runes[i])

		} // end inner for

		if indx == len(words)-1 {
			resultString += newString
		} else {
			if countSpaces > 1 {
				newSpaces := ""
				for i := 0; i < countSpaces; i++ {
					newSpaces += " "
				} // end inner for
				resultString += newString + newSpaces
			} else {
				resultString += newString + " "
			} // end inner if
		} // end 2 inner for

	} // end outer for

	fmt.Println(countSpaces)
	fmt.Println(resultString == "elbuod")
}


// !! ANSWER

/* func ReverseWords(str string) string {
	var rev string
	var word string
	
	for _, i := range str {
	  if i == ' ' {
		rev = rev + word + " " // Adds word and space to result
		word = "" // Empties word variable
	  } else {
		word = string(i) + word // Adds letter to temporary word variable
	  } 
	}
	
	return rev + word// reverse those words
	} */