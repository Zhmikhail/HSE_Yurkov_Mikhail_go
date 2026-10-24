package main

import (
	"fmt"
	"math/big"
	"mentor-seminar/input"
)

func main() {
	fmt.Print("Введите N: ")
	n, err := input.ScanInt()
	if err != nil || n <= 0 {
		fmt.Println("Ошибка: N должно быть положительным целым числом")
		return
	}

	// Формула в big.Int сохраняет точность даже для очень большого N.
	bigN := big.NewInt(int64(n))
	sum := new(big.Int).Mul(bigN, new(big.Int).Add(bigN, big.NewInt(1)))
	sum.Div(sum, big.NewInt(2))
	average, _ := new(big.Float).Quo(new(big.Float).SetInt(sum), new(big.Float).SetInt(bigN)).Float64()
	fmt.Printf("Сумма: %s\nСреднее: %g\n", sum.String(), average)
}
