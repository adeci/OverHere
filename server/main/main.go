package main

import (
	"OverHere/server/greetings"
	"OverHere/server/handler"
)

func main() {
	greetings.Greetings()
	greetings.Greetings()

	//execCommands.Run()
	handler.Handle()
}
