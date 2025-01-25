package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ServerIP := "192.168.10.191"
	Port := "9090"
	ServerAddress := ServerIP + ":" + Port

	connection, err := net.Dial("tcp", ServerAddress)
	if err != nil {
		fmt.Println("Connection couldnt be established with tthe server")
		log.Fatal(err)
	}

	fmt.Println("[+] Connection established with: ", connection.RemoteAddr().String())
}
