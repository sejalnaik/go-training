package main

import (
	"errors"
	"fmt"
)

func main() {
	qoutient, err := divide(10.0, 0.0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f\n", qoutient)
	}

}

func divide(dividend float64, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0, errors.New("Cant divide by zero")
	}
	return dividend / divisor, nil
}
