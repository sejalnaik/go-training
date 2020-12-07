package main

import (
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

type myError struct {
	s string
}

func (err *myError) Error() string {
	return err.s
}

func divide(numerator int, denominator int) (int, error) {
	if denominator == 0 {
		return 0, &myError{s: "Cannot divide by 0"}
	}
	return numerator / denominator, nil
}
