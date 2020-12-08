package pkginit

import "fmt"

var num int

func init() {
	num = 40
	fmt.Println(num)
}
