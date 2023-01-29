package Project1

import "fmt"

func main() {
	var numbers int
	maxNumber := 0
	for c := 0; c < 3; c++ {
		fmt.Scan(&numbers)
		if numbers > maxNumber {
			maxNumber = numbers
		}
	}
	fmt.Println(maxNumber)
}
