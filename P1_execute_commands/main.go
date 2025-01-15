package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	command := "date"

	cmd_obj := exec.Command(command)
	cmd_obj.Stdout = os.Stdout
	cmd_obj.Stderr = os.Stderr

	err := cmd_obj.Run()

	if err != nil {
		log.Fatal(err)
	}
}
