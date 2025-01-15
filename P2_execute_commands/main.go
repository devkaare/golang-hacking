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
	cmd_obj.Stdin = os.Stdin

	err = cmd_obj.Run()

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func main() {
	executeCommand("sudo", []string{"ifconfig", "eth0", "down"})
	executeCommand("sudo", []string{"ifconfig", "eth0", "hw", "ether", "00:11:22:33:44:55"}) // fe80::fae3:8aae:8ab5:b4a3
	executeCommand("sudo", []string{"ifconfig", "eth0", "up"})
}
