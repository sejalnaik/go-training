package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("myFile.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Content:", string(content))
}
