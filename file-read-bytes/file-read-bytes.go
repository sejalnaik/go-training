package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("myFile.txt")
	defer f.Close()
	checkError(err)
	bucket := make([]byte, 5)
	bytesRead, er := f.Read(bucket)
	checkError(er)
	fmt.Println("Content(limited to size of bucket):", string(bucket), " Number of bytes read:", bytesRead)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
