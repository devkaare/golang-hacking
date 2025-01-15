// NOTE: Don't know why the course creator used snake case, but I'm following along so I will aswell... FOR NOW

package main

import (
	"log"
	"os"
	"os/exec"
)

func executeCommand(command string, args_arr []string) (err error) {
	args := args_arr
	cmd_obj := exec.Command(command, args...)
	cmd_obj.Stdout = os.Stdout
	cmd_obj.Stderr = os.Stderr

	err = cmd_obj.Run()

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func main() {
	command := "ls"
	executeCommand(command, []string{"-l"})
}
