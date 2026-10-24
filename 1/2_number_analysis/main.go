package main

import (
	"fmt"
	"mentor-seminar/input"
)

func main() {
	fmt.Print("Введите число N: ")
	n, err := input.ScanInt()
	if err != nil {
		fmt.Println("Ошибка: требуется целое число")
		return
	}

	if n > 0 {
		fmt.Println("Число положительное")
	} else if n < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Число равно нулю")
	}
	if n%2 == 0 {
		fmt.Println("Число чётное")
	} else {
		fmt.Println("Число нечётное")
	}
	if n%3 == 0 && n%5 == 0 {
		fmt.Println("Число делится и на 3, и на 5")
	} else {
		fmt.Println("Число не делится одновременно на 3 и на 5")
	}
}
