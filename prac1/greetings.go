package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("No name")
	}
	message := fmt.Sprintf("Welcome, %v!", name)
	return message, nil
}
