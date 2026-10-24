package main

import (
	"fmt"
	"mentor-seminar/input"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for divisor := 2; divisor <= n/divisor; divisor++ {
		if n%divisor == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Print("Введите число: ")
	n, err := input.ScanInt()
	if err != nil {
		fmt.Println("Ошибка: требуется целое число")
		return
	}
	if isPrime(n) {
		fmt.Printf("Число %d простое\n", n)
	} else {
		fmt.Printf("Число %d не является простым\n", n)
	}
}
