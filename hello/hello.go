package main

import (
	"fmt"
	"log"

	greetings "example.com/prac1"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	msg, err := greetings.Hello("Raymond")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
	msg, err = greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}
