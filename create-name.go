package main

import (
	"fmt"
	"strings"
)

func main() {
	fullName := createName("sejal", "naik")
	fmt.Println(fullName)
}

func createName(firstName string, lastName string) string {
	fullName := strings.ToUpper(firstName + " " + lastName)
	return fullName
}
