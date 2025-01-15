package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	command := "date"

	// NOTE: Don't know why the course creator used snake case, but I'm following along so I will aswell... FOR NOW
	cmd_obj := exec.Command(command)
	cmd_obj.Stdout = os.Stdout
	cmd_obj.Stderr = os.Stderr

	err := cmd_obj.Run()

	if err != nil {
		log.Fatal(err)
	}
}
