package main

import (
	"fmt"
	"mentor-seminar/input"
)

func main() {
	fmt.Print("Введите количество чисел: ")
	n, err := input.ScanInt()
	if err != nil || n < 0 {
		fmt.Println("Ошибка: количество не может быть отрицательным")
		return
	}

	positive, negative, zero := 0, 0, 0
	for i := 0; i < n; i++ {
		value, err := input.ScanInt()
		if err != nil {
			fmt.Println("Ошибка: требуется целое число")
			return
		}
		if value > 0 {
			positive++
		} else if value < 0 {
			negative++
		} else {
			zero++
		}
	}
	fmt.Printf("Положительных: %d\nОтрицательных: %d\nНулей: %d\n", positive, negative, zero)
}
