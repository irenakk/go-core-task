package main

import (
	"fmt"
)

func CubePipeline(input <-chan uint8, output chan<- float64) {
	for num := range input {
		cubed := float64(num) * float64(num) * float64(num)
		output <- cubed
	}
	close(output)
}

// Сделать конвейер чисел
// Даны два канала.
// В первый пишутся числа типа uint8. Нужно, чтобы числа читались из первого канала по мере поступления,
// затем эти числа должны преобразовываться в float64 и возводиться в куб и результат записывался во второй канал.
// Напишите main функцию, в которой протестируете весь вышеописанный функционал. Выведите результаты на экран.
func main() {
	input := make(chan uint8)
	output := make(chan float64)

	go func() {
		for x := uint8(1); x <= 5; x++ {
			input <- x
		}
		close(input)
	}()

	go CubePipeline(input, output)

	for result := range output {
		fmt.Printf("%.0f\n", result)
	}
}
