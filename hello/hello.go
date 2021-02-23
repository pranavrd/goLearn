package main

import (
	"fmt"
	"log"

	greetings "example.com/prac1"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Raymond", "Jacob", "Amy"}

	msg, err := greetings.Hello(names)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range msg {
		fmt.Printf("%v\n", v)
	}
	//fmt.Println(msg)
}
