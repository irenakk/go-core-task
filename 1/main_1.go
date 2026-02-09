package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// HashRuneSliceWithSalt принимает срез рун, добавляет соль в середину и возвращает SHA256 хэш
func HashRuneSliceWithSalt(runes []rune, salt string) string {
	mid := len(runes) / 2
	firstHalf := runes[:mid]
	secondHalf := runes[mid:]
	saltedString := string(firstHalf) + salt + string(secondHalf)
	hash := sha256.Sum256([]byte(saltedString))
	return hex.EncodeToString(hash[:])
}

func main() {
	// 1. Создает несколько переменных различных типов данных:
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	// 2. Определяет тип каждой переменной и выводит его на экран.
	fmt.Printf("%T\n", numDecimal)
	fmt.Printf("%T\n", numOctal)
	fmt.Printf("%T\n", numHexadecimal)
	fmt.Printf("%T\n", pi)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", isActive)
	fmt.Printf("%T\n", complexNum)

	// 3. Преобразует все переменные в строковый тип и объединяет их в одну строку.
	var bigString = fmt.Sprint(numDecimal) + fmt.Sprint(numOctal) + fmt.Sprint(numHexadecimal) + fmt.Sprint(pi) + fmt.Sprint(name) + fmt.Sprint(isActive) + fmt.Sprint(complexNum)
	fmt.Println(bigString)

	// 4. Преобразовать эту строку в срез рун.
	runeSlice := []rune(bigString)
	fmt.Println(runeSlice)
	fmt.Printf("%U\n", runeSlice)

	// 5. Захэшировать этот срез рун SHA256, добавив в середину соль "go-2024" и вывести результат.
	salt := "go-2024"
	hashResult := HashRuneSliceWithSalt(runeSlice, salt)
	fmt.Println("SHA256 хэш с солью:", hashResult)
}
