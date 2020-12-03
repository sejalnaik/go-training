package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Hour()

	if now <= 20 {
		fmt.Println("Good morning")
	} else if now <= 17 {
		fmt.Println("Good evening")
	} else {
		fmt.Println("Good night")
	}

}
