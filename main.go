package main

import "fmt"

// Функция для генерации всех перестановок элементов в массиве
func generatePermutations(arr []int, perm []int, used []bool) {
	// Базовый случай: если длина текущей перестановки равна длине исходного массива,
	// то выводим текущую перестановку
	if len(perm) == len(arr) {
		fmt.Println(perm)
		return
	}

	// Рекурсивно генерируем все возможные перестановки
	for i := 0; i < len(arr); i++ {
		if !used[i] {
			// Помечаем элемент как использованный
			used[i] = true
			// Добавляем текущий элемент в перестановку
			perm = append(perm, arr[i])
			// Рекурсивно генерируем остальные перестановки
			generatePermutations(arr, perm, used)
			// Удаляем последний элемент, чтобы подготовиться к следующей итерации
			perm = perm[:len(perm)-1]
			// Снимаем пометку о использовании текущего элемента
			used[i] = false
		}
	}
}

func main() {
	arr := []int{1, 2, 3}
	perm := make([]int, 0, len(arr))
	used := make([]bool, len(arr))

	// Генерируем все перестановки
	generatePermutations(arr, perm, used)
}
