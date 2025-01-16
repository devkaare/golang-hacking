package main

import (
	"flag"
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
	iface := flag.String("iface", "", "Interface for which you want to change the MAC")
	newMac := flag.String("newMac", "", "provide the new mac address")

	flag.Parse()

	executeCommand("sudo", []string{"ifconfig", *iface, "down"})
	executeCommand("sudo", []string{"ifconfig", *iface, "hw", "ether", *newMac})
	executeCommand("sudo", []string{"ifconfig", *iface, "up"})
}
