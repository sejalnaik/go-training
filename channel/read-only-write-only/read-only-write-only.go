package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	ch := make(chan int, 2)
	go writeToChannel(ch)
	go readFromChannel(ch)
	wg.Wait()
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
