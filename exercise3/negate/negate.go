package main

import "fmt"

func main() {
	truth := true
	negate(&truth)
	fmt.Println(truth)

	lies := false
	negate(&lies)
	fmt.Println(lies)
}

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}
