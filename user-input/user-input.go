package main

import (
	"fmt"
	"log"
)

func main() {
	//input using bufio
	/*fmt.Println("Enter a word")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered :", input)*/

	//input using scan
	fmt.Println("Enter a number")
	var number int
	//the type of input expected is the type of the variable given to fmt.Scan()
	_, err := fmt.Scan(&number)
	if err != nil {
		log.Fatal("Error occurred:", err)
	}
	fmt.Println("You entered :", number, "Error:", err)
}
