// So Many Permutations - 4 kyu
// https://www.codewars.com/kata/5254ca2719453dcc0b00027d/train/go
// test#17.4+ready

package main

import (
	"fmt"
	//"strings"
)

func main() {
	startStr := "aabb"
	var arrLiteral []string
	var resulatArr []string

	usedVar := make(map[string]bool)

	for _, str := range startStr {
		arrLiteral = append(arrLiteral, string(str))
	}

	vars := make([]string, 0, len(arrLiteral))
	used := make([]bool, len(arrLiteral))

	generateVar(arrLiteral, vars, used, &resulatArr, usedVar)

	fmt.Println(resulatArr)
}

func generateVar(arr []string, vars []string, used []bool, resulatArr *[]string, usedVar map[string]bool) {
	if len(vars) == len(arr) {
		newStr := ""
		for _, s := range vars {
			newStr += s
		}
		if !usedVar[newStr] {
			usedVar[newStr] = true
			*resulatArr = append(*resulatArr, newStr)
			return
		}
		return
	}

	for i := 0; i < len(arr); i++ {
		if !used[i] {
			used[i] = true
			vars = append(vars, arr[i])
			generateVar(arr, vars, used, resulatArr, usedVar)
			vars = vars[:len(vars)-1]
			used[i] = false
		}
	}
}


// !! BEST ANSWER

/* func Permutations(s string) (a[]string) {
  if len(s) < 2 { return []string{s} }
  m := map[string]bool{}
  for _, sub := range Permutations(s[1:]) {
    for i := 0; i <= len(sub); i++ {
      st := sub[0:i] + s[0:1] + sub[i:]
      if m[st] { continue }
      m[st]=true
      a = append(a, st)
    }
  }
  return
} */

