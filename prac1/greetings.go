package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(names []string) (map[string]string, error) {

	msgs := make(map[string]string)
	for _, name := range names {
		msg, err := func(name string) (string, error) {
			if name == "" {
				return name, errors.New("No name")
			}
			msg := fmt.Sprintf(randomFormat(), name)
			return msg, nil
		}(string(name))

		if err != nil {
			return nil, err
		}
		msgs[string(name)] = msg
	}
	return msgs, nil
}

func HelloOld(name string) (string, error) {
	if name == "" {
		return name, errors.New("No name")
	}
	msg := fmt.Sprintf(randomFormat(), name)
	return msg, nil
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
