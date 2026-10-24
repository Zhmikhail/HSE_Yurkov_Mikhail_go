package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	var operation string

	fmt.Print("Введите первое число: ")
	if _, err := fmt.Scan(&a); err != nil {
		fmt.Println("Ошибка: первое значение должно быть числом")
		return
	}
	fmt.Print("Введите операцию (+, -, *, /): ")
	fmt.Scan(&operation)
	fmt.Print("Введите второе число: ")
	if _, err := fmt.Scan(&b); err != nil {
		fmt.Println("Ошибка: второе значение должно быть числом")
		return
	}

	var result float64
	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль невозможно")
			return
		}
		result = a / b
	default:
		fmt.Println("Ошибка: неизвестная операция")
		return
	}
	if math.IsNaN(result) || math.IsInf(result, 0) {
		fmt.Println("Ошибка: результат не помещается в число с плавающей точкой")
		return
	}
	fmt.Printf("Результат: %g\n", result)
}
