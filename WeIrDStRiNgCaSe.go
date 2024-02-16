// WeIrD StRiNg CaSe - 6 kyu
// teset#9.1

package main

import (
	"fmt"
	//"strings"
	"unicode"
)

func main() {
	str := "This is a test Looks like you passed"
	i := 0

	var newString []rune
	for _, char := range str {
		if char == ' ' {
			i = 0
			newString = append(newString, char)
			continue
		}

		if i == 0 || i%2 == 0 {
			newString = append(newString, unicode.ToUpper(char))
		} else {
			newString = append(newString, unicode.ToLower(char))
		}
		i++
	}

	fmt.Println(string(newString) == "ThIs Is A TeSt LoOkS LiKe YoU PaSsEd")
}
