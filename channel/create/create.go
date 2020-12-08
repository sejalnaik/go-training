package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start main")
	ch := make(chan int)
	go printChannelValues(ch)
	//waits for printChannelValues() after giving data to channel
	ch <- 10
	fmt.Println("value of ch after changed by func :", <-ch)
	close(ch)
	fmt.Println("End main")
}

func printChannelValues(ch chan int) {
	var chanReceivedNumber int = <-ch
	fmt.Println("Chan value recieved:", chanReceivedNumber)
	var chanVlueChanged = 20
	ch <- chanVlueChanged
}
