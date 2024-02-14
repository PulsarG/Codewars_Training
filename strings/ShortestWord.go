// Shortest Word

package main

import (
	"fmt"
	"strings"
)

func main() {
	longString := "dfg dfgg dfggg sf"

	words := strings.Fields(longString)

	length := 1

	for i, word := range words {
		if i == 0 {
			length = len(word)
		}

		if length > len(word) {
			length = len(word)
		}
	}

	fmt.Println(length)
}


// !! ANSWER

/* func FindShort(s string) int {
	shortest := len(s)
	for _, word := range strings.Split(s, " ") {
	  if len(word) < shortest {
		shortest = len(word)
	  }
	}
	return shortest
  } */
