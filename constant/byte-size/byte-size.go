package main

import "fmt"

type byteSize float64

const (
	_           = iota
	KB byteSize = 1 << (10 * iota)
	MB
	GB
	TB
)

func main() {
	var fileSize byteSize = 50000000.0
	fmt.Printf("%.2fKB", fileSize/KB)
}
