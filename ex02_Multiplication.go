package main

import "fmt"

func main() {
	var numbers int
	var maxNumber int = 0
	var multiplication int = 1
	var amount int
	fmt.Print("Введите желаемое количество чисел: ")
	fmt.Scan(&amount)
	fmt.Print("Введите свои числа: ")

	for c := 0; c < amount; c++ {
		fmt.Scan(&numbers) // Введите свои числа
		if numbers > maxNumber {
			maxNumber = numbers
		}
	}

	for c := 1; c <= maxNumber; c++ {
		multiplication *= c
	}

	fmt.Print(multiplication) // Ответ
}
