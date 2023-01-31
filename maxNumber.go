package main

import "fmt"

func main() {
	var numbers int
	var maxNumber int = 0
	for c := 0; c < 3; c++ {
		fmt.Scan(&numbers) // Введите свои числа
		if numbers > maxNumber {
			maxNumber = numbers
		}
	}
	fmt.Println(maxNumber) // Вывод максимального числа
}
