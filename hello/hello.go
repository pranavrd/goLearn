package main

import (
	"fmt"

	greetings "example.com/prac1"
)

func main() {
	msg := greetings.Hello("Raymond")
	fmt.Println(msg)
}
