package main

import "fmt"

// Напишите функцию, которая проверяет, есть ли пересечения значений между двумя слайсами и возвращает:
// - bool значение есть ли хотя бы одно пересечение в значениях входных срезов
// - срез []int с пересеченными значениями (если таких значений нет, то возвращать пустой срез).
func uniqueElements(slice1, slice2 []int) (bool, []int) {
	elementsToRemove := make(map[int]struct{})
	for _, item := range slice2 {
		elementsToRemove[item] = struct{}{}
	}

	var result []int
	for _, item := range slice1 {
		if _, found := elementsToRemove[item]; found {
			result = append(result, item)
		}
	}

	var intersection bool
	if len(result) != 0 {
		intersection = true
	} else {
		intersection = false
	}

	return intersection, result
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	fmt.Println(uniqueElements(a, b))
}
