package main

import (
	"errors"
	"fmt"
)

func main() {
	quotient, err := divide(4, 0)
	if err != nil {
		fmt.Println("error occurred:", err)
		return
	}
	fmt.Println("Result is:", quotient)

}

func divide(numerator int, denominator int) (int, error) {
	if denominator == 0 {
		return 0, errors.New("Cannot divide by 0")
	}
	return numerator / denominator, nil
}
