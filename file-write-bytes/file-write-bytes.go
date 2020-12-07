package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("myFile.txt")
	checkError(err)
	bytesWritten, er := f.WriteString("i am sejal naik")
	checkError(er)
	fmt.Println("Bytes written:", bytesWritten)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
