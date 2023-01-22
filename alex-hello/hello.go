package main

import (
	"fmt"
	"log"

	greetings "github.com/adeci/OverHere/alex-greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Billy Bob Jenkins")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
