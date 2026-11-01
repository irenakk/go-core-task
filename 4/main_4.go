package main

import "fmt"

// Напишите функцию, которая возвращает слайс строк, содержащий элементы, которые есть в первом слайсе, но отсутствуют во втором.
func uniqueElements(slice1, slice2 []string) []string {
	elementsToRemove := make(map[string]struct{})
	for _, item := range slice2 {
		elementsToRemove[item] = struct{}{}
	}

	var result []string
	for _, item := range slice1 {
		if _, found := elementsToRemove[item]; !found {
			result = append(result, item)
		}
	}

	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	fmt.Println(uniqueElements(slice1, slice2))
}
