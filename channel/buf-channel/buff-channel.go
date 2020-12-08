package main

import "fmt"

func main() {
	fmt.Println("Start of main")
	messageChannel := make(chan string, 3)
	go printMessages(messageChannel)
	// does not wait for printMessages() after giving data to channel
	messageChannel <- "first message"
	messageChannel <- "second message"
	messageChannel <- "third message"
	messageChannel <- "fourth message"
	fmt.Println("End of main")
}

func printMessages(messageChannel chan string) {
	var message string
	message = <-messageChannel
	message = <-messageChannel
	//message = <-messageChannel
	//message = <-messageChannel
	//message = <-messageChannel
	fmt.Println(message)
}
