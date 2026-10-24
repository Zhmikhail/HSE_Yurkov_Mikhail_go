package main

import (
	"fmt"
	"math/big"
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
	sum := new(big.Int)
	even, odd := 0, 0
	for i := 0; i < n; i++ {
		value, err := input.ScanInt()
		if err != nil {
			fmt.Println("Ошибка: требуется целое число")
			return
		}
		sum.Add(sum, big.NewInt(int64(value)))
		if i == 0 || value < minimum {
			minimum = value
		}
		if i == 0 || value > maximum {
			maximum = value
		}
		if value%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	average, _ := new(big.Float).Quo(new(big.Float).SetInt(sum), big.NewFloat(float64(n))).Float64()
	fmt.Printf("Сумма: %s\nСреднее: %g\nМаксимум: %d\nМинимум: %d\nЧётных: %d\nНечётных: %d\n",
		sum.String(), average, maximum, minimum, even, odd)
}
