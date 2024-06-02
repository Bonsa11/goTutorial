package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("greeting: ")
	log.SetFlags(log.Ldate | log.Ltime)

	// Get a greeting message and print it.
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
