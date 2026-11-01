package main

import (
	"fmt"
	"math/rand"
)

// Напишите генератор случайных чисел используя небуфферизированный канал.
func RandomGenerator(ch chan<- int, count int) {
	for i := 0; i < count; i++ {
		num := rand.Intn(100)
		ch <- num
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go RandomGenerator(ch, 10)

	for num := range ch {
		fmt.Printf("Получено: %d\n", num)
	}
}
