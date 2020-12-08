package main

import (
	"fmt"
	"time"
)

func main() {
	go DelayOne()
	go DelayTwo()
	fmt.Scanln()
}

func DelayOne() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("DelayOne Count :", i+1)
	}
	fmt.Println("End of DelayOne")
}

func DelayTwo() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("DelayTwo Count :", i+1)
	}
	fmt.Println("End of DelayTwo")
}
