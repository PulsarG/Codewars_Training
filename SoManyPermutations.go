// So Many Permutations - 4 kyu
// https://www.codewars.com/kata/5254ca2719453dcc0b00027d/train/go
// test#17.1.1

package main

import (
	"fmt"
	//"strings"
)

func main() {
	startStr := "abc"
	var arrLiteral []string
	var resulatArr []string
	ignotInd := 4

	for _, str := range startStr {
		arrLiteral = append(arrLiteral, string(str))
	}

	generateCombinations(arrLiteral, resulatArr, ignotInd)

	fmt.Println(arrLiteral)
}

func generateCombinations(arrLit, resulatArr []string, ignotInd int) {
	newStr := ""
	for i := 0; i < len(arrLit); i++ {
		newStr += arrLit[i]
	}
}

// ** Первый индекс + следующий + следующий+1...
// ** Первый индекс + следущий+1 + следующий + следующий+2...

//В цикле индекс = длина - 1
//	индекс литерал - первый в новой строке
//	запуск метода перебора всех литералов, кроме индекса
//		внутри функции рекурсия - создается массив без индекса и цикл со старта

//generateCombinations("", n)

// Разбиение на руны
// Алгорит перебор
// сбор в строки
// Добавление в массив

/* With input 'a':
Your function should return: ['a']

With input 'ab':
Your function should return ['ab', 'ba']

With input 'abc':
Your function should return ['abc','acb','bac','bca','cab','cba'] */
