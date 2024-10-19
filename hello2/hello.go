package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set Properties of the predefied Logger, including
	// the log entry prefix and a flag to disable priting
	// the time, source file and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names
	names := []string{"John", "Peter", "Steven"}

	names = append(names, "Chris")

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(messages)
}
