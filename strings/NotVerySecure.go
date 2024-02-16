// Not very secure - 5 kyu
// test#11.2+ready

package main

import (
	"fmt"
	//"strconv"
	//"strings"
	//"unicode"
	"regexp"
)

func main() {
	var IsLetter = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString

	fmt.Println(IsLetter("golang"))
	fmt.Println(IsLetter("123"))
	fmt.Println(IsLetter("рои134432-=-!№;;"))
}


// !! BEST ANSWER 

/* func alphanumeric(s string) bool {
	r := regexp.MustCompile("^[a-zA-Z0-9]+$")
	return r.MatchString(s)
} */
