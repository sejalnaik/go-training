package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Close resource")
		details := recover()
		fmt.Println("Details:", details)
	}()
	fmt.Println("Open resource")
	panic("I dont know what to do")
	fmt.Println("End of main")
}
