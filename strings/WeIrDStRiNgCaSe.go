// WeIrD StRiNg CaSe - 6 kyu
// teset#9.2+reasy

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


// !! BEST ANSWER

/* func toWeirdCase(str string) string {
	cap := true 
	var result []rune
	for _, c := range str {
	  if(c == 32){
		result = append(result, c)
		cap = true
	  }else if(cap){
		result = append(result, unicode.ToUpper(c))
		cap = false
	  }else{
		result = append(result, unicode.ToLower(c))
		cap = true
	  }
	}
	return string(result)
  } */
