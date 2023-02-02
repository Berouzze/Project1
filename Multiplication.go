package main

import "fmt"

func main() {
	var numbers int
	var maxNumber int = 0
	var multiplication int = 1

	for c := 0; c < 3; c++ {
		fmt.Scan("Введите 3 своих числа", &numbers) // Введите свои числа
		if numbers > maxNumber {
			maxNumber = numbers
		}
	}

	for c := 1; c <= maxNumber; c++ {
		multiplication *= c
	}

	fmt.Print("Ответ ", multiplication) // Ответ
}
