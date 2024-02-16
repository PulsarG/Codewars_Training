// Not very secure - 5 kyu
// test#11.1

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
