package main

import (
	"fmt"
	"math/big"
	"mentor-seminar/input"
)

func main() {
	fmt.Print("Введите N: ")
	n, err := input.ScanInt()
	if err != nil || n < 1 {
		fmt.Println("Ошибка: N должно быть положительным целым числом")
		return
	}

	// Подходящие числа — это все кратные 6. Формула не переполняется.
	k := big.NewInt(int64(n / 6))
	sum := new(big.Int).Mul(k, new(big.Int).Add(k, big.NewInt(1)))
	sum.Mul(sum, big.NewInt(3))
	fmt.Printf("Сумма подходящих чисел: %s\n", sum.String())
}
