package main

import (
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	ch := make(chan struct{})
	wg.Wait()
	go func() {
		ch <- struct{}{}
	}()
}
