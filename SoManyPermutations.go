// So Many Permutations - 4 kyu
// https://www.codewars.com/kata/5254ca2719453dcc0b00027d/train/go
// test#17.2

package main

import (
	"fmt"
	//"strings"
)

func main() {
	startStr := "abc"
	var arrLiteral []string

	for _, str := range startStr {
		arrLiteral = append(arrLiteral, string(str))
	}


	fmt.Println(arrLiteral)
}

В цикле индекс = длина - 1
	индекс литерал - первый в новой строке
	запуск метода перебора всех литералов, кроме индекса
		внутри функции рекурсия - создается массив без индекса и цикл со старта

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
