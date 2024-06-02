package greetings

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (message string, err error) {
	if name == "" {
		return "", errors.New("no name was provided")
	}

	// Return a greeting that embeds the name in a message.
	message = fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (messages map[string]string, err error) {
	messages = make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

func randomFormat() (message string) {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
