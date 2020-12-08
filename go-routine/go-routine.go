package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup = sync.WaitGroup{}

func main() {
	wg.Add(0)
	go delayOne()
	go delayTwo()
	wg.Wait()
	fmt.Println("End of main")
}

func delayOne() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("DelayOne Count :", i+1)
	}
	fmt.Println("End of DelayOne")
	wg.Done()
}

func delayTwo() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("DelayTwo Count :", i+1)
	}
	fmt.Println("End of DelayTwo")
	wg.Done()
}
