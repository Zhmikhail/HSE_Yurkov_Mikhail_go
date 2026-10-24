package main

import (
	"fmt"
	"math/big"
	"mentor-seminar/input"
)

func printTable(n int) {
	bigN := big.NewInt(int64(n))
	for i := 1; i <= 10; i++ {
		result := new(big.Int).Mul(bigN, big.NewInt(int64(i)))
		fmt.Printf("%d * %d = %s\n", n, i, result.String())
	}
}

func main() {
	fmt.Print("Введите число: ")
	n, err := input.ScanInt()
	if err != nil {
		fmt.Println("Ошибка: требуется целое число")
		return
	}
	printTable(n)
}
