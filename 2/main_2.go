package main

import (
	"fmt"
	"math/rand"
)

// 1. Создайте слайс целых чисел originalSlice, содержащий 10 произвольных значений, которые генерируются случайным
// образом (при каждом запуске должны получаться новые значения)
func randSlice(len int) []int {
	slice := make([]int, len)
	for i := 0; i <= len-1; i++ {
		slice[i] = rand.Intn(100)
	}
	return slice
}

// 2. Напишите функцию sliceExample, которая принимает слайс и возвращает новый слайс, содержащий только четные числа из исходного слайса.
func sliceExample(oldSlice []int) []int {
	var newSlice []int
	for i := 0; i < len(oldSlice); i++ {
		if oldSlice[i]%2 == 0 {
			newSlice = append(newSlice, oldSlice[i])
		}
	}
	return newSlice
}

// 3. Напишите функцию addElements, которая принимает слайс и число. Функция должна добавлять это число в конец слайса и возвращать новый слайс.
func addElements(oldSlice []int, num int) []int {
	var newSlice []int
	newSlice = append(oldSlice, num)
	return newSlice
}

// 4. Напишите функцию copySlice, которая принимает слайс и возвращает его копию. Убедитесь, что изменения в оригинальном слайсе не влияют на его копию.
func copySlice(oldSlice []int) []int {
	copySlice := make([]int, len(oldSlice))
	copy(copySlice, oldSlice)
	return copySlice
}

// 5. Напишите функцию removeElement, которая принимает слайс и индекс элемента, который нужно удалить. Функция должна возвращать новый слайс без элемента по указанному индексу.
func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	result := make([]int, 0, len(slice)-1)
	result = append(result, slice[:index]...)
	result = append(result, slice[index+1:]...)
	return result
}

// 6. Напишите main функцию, в которой протестируете все вышеописанные функции. Выведите результаты на экран.
func main() {
	// 1. Создайте слайс целых чисел originalSlice, содержащий 10 произвольных значений, которые генерируются случайным
	// образом (при каждом запуске должны получаться новые значения)
	length := 10
	originalSlice := randSlice(length)
	fmt.Println(originalSlice)

	// 2. Напишите функцию sliceExample, которая принимает слайс и возвращает новый слайс, содержащий только четные числа из исходного слайса.
	evenSlice := sliceExample(originalSlice)
	fmt.Println(evenSlice)

	// 3. Напишите функцию addElements, которая принимает слайс и число. Функция должна добавлять это число в конец слайса и возвращать новый слайс.
	num := 9
	extendedSlice := addElements(evenSlice, num)
	fmt.Println(extendedSlice)

	// 4. Напишите функцию copySlice, которая принимает слайс и возвращает его копию. Убедитесь, что изменения в оригинальном слайсе не влияют на его копию.
	copiedSlice := copySlice(originalSlice)
	fmt.Println(copiedSlice)
	originalSlice[0] = 999
	fmt.Println(originalSlice)
	fmt.Println(copiedSlice)

	// 5. Напишите функцию removeElement, которая принимает слайс и индекс элемента, который нужно удалить. Функция должна возвращать новый слайс без элемента по указанному индексу.
	testSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(testSlice)
	removedSlice := removeElement(testSlice, 2)
	fmt.Println(removedSlice)
}
