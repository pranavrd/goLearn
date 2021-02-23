package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Welcome, %v!", name)
	return message
}
