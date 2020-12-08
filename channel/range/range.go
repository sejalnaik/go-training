package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	ch := make(chan int, 2)
	ch2 := make(chan struct{})
	go writeToChannel(ch)
	go readFromChannel(ch)
	wg.Wait()
	close(ch)
	go func() {
		ch2 <- struct{}{}
	}()

	for i := range ch {
		fmt.Println("Inside range of ch:", i)
	}

	for i := 0; i < 2; i++ {
		select {
		case chan1, ok := <-ch:
			fmt.Println("Channel 1 has data", chan1, ok)
		case chan2, ok := <-ch2:
			fmt.Println("Channel 2 has data", chan2, ok)

		}
	}
}

func writeToChannel(ch chan<- int) {
	fmt.Println("Starting to write")
	ch <- 10
	ch <- 5
	ch <- 0
	fmt.Println("Inside writeToChannel(), data has been written to channel")
	wg.Done()
}

func readFromChannel(ch <-chan int) {
	fmt.Println("Starting to read")
	value1 := <-ch
	value2 := <-ch
	fmt.Println("Inside readFromChannel(), data has been read from channel", value1, value2)
	wg.Done()
}
