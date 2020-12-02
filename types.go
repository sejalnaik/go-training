package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf("Hello world"))
	fmt.Println(reflect.TypeOf(10))
	fmt.Println(reflect.TypeOf(10.1))
	fmt.Println(reflect.TypeOf(true))
}
