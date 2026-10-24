package main

import (
	"fmt"
	"mentor-seminar/input"
)

func main() {
	fmt.Print("Введите количество чисел: ")
	n, err := input.ScanInt()
	if err != nil || n <= 0 {
		fmt.Println("Ошибка: количество должно быть положительным целым числом")
		return
	}

	var minimum, maximum int
	for i := 0; i < n; i++ {
		value, err := input.ScanInt()
		if err != nil {
			fmt.Println("Ошибка: требуется целое число")
			return
		}
		if i == 0 || value < minimum {
			minimum = value
		}
		if i == 0 || value > maximum {
			maximum = value
		}
	}
	fmt.Printf("Минимум: %d\nМаксимум: %d\n", minimum, maximum)
}
