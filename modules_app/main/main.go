package main

import (
	"modules_app/greeting"
	"modules_app/greeting/hindi"

	greet "github.com/headfirstgo/greeting"
)

func main() {
	greeting.SayHello()
	hindi.SayHello()
	greet.Hello()
}
