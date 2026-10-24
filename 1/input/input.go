package input

import (
	"fmt"
	"strconv"
)

// ScanInt читает ровно один токен и не принимает дробные значения вроде 1.5.
func ScanInt() (int, error) {
	var token string
	if _, err := fmt.Scan(&token); err != nil {
		return 0, err
	}
	return strconv.Atoi(token)
}
