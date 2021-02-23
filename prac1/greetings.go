package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("No name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	f := []string{
		"Welcome, %v", "Hello, %v", "Nice to meet you, %v",
	}

	return f[rand.Intn(len(f))]
}
