package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data := []byte("I a have written a sentence")
	err := ioutil.WriteFile("myFile.txt", data, 0644)
	checkError(err)
	fmt.Println("done writing")
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
