package main

import "fmt"

func main() {

	//LIFO order
	var num = 50
	defer f1()
	// the value of num is 50 since its argument value is fixed in the statement itself
	defer f2(num)
	num = 100
	fmt.Println("End of main. Number :", num)
}

func f1() {
	defer fmt.Println("Hello form f1")
	fmt.Println("End of f1")
}

func f2(num int) {
	fmt.Println("End of f2. Number :", num)
}
