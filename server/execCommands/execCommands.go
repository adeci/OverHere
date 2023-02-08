package execCommands

import (
	"log"
	"os/exec"
)

func Run() {

	cmd := exec.Command("firefox")

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
