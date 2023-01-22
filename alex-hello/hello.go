package main

import (
	"fmt"

	greetings "github.com/adeci/OverHere/alex-greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
