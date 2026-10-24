package main

import (
	"fmt"
	"math"
	"mentor-seminar/input"
)

func isFinite(value float64) bool {
	return !math.IsNaN(value) && !math.IsInf(value, 0)
}

func main() {
	var balance float64
	fmt.Print("Введите начальный баланс: ")
	if _, err := fmt.Scan(&balance); err != nil || !isFinite(balance) || balance < 0 {
		fmt.Println("Ошибка: баланс должен быть неотрицательным числом")
		return
	}

	fmt.Println("1 — посмотреть баланс")
	fmt.Println("2 — пополнить баланс")
	fmt.Println("3 — снять средства")
	fmt.Print("Выберите действие: ")
	action, err := input.ScanInt()
	if err != nil {
		fmt.Println("Ошибка: требуется номер действия")
		return
	}

	switch action {
	case 1:
		fmt.Printf("Баланс: %.2f\n", balance)
	case 2, 3:
		var amount float64
		fmt.Print("Введите сумму: ")
		if _, err := fmt.Scan(&amount); err != nil || !isFinite(amount) || amount <= 0 {
			fmt.Println("Ошибка: сумма должна быть положительной")
			return
		}
		if action == 2 {
			balance += amount
			if !isFinite(balance) {
				fmt.Println("Ошибка: баланс не помещается в число с плавающей точкой")
				return
			}
			fmt.Printf("Баланс после пополнения: %.2f\n", balance)
		} else if amount > balance {
			fmt.Println("Ошибка: недостаточно средств")
		} else {
			balance -= amount
			fmt.Printf("Баланс после снятия: %.2f\n", balance)
		}
	default:
		fmt.Println("Ошибка: неизвестное действие")
	}
}
