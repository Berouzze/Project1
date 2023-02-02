package main

import "fmt"

func main() {
	var numbers, maxNumber int

	var multiplication int = 1

	fmt.Print("Введите своих числа: ")

	for c := 0; c < 3; c++ {
		fmt.Scan(&numbers) // Введите свои числа
		if numbers > maxNumber {
			maxNumber = numbers
		}
	}

	for c := 1; c <= maxNumber; c++ {
		multiplication *= c
	}

	fmt.Print("Ответ ", multiplication) // Ответ на задачу
}
