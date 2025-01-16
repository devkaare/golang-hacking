package main

import (
	"fmt"
	"net"
)

func main() {
	IP := "scanme.nmap.org"
	Port := "80"

	address := IP + ":" + Port
	connection, err := net.Dial("tcp", address)

	if err == nil {
		fmt.Println("[+] Connection established..", connection.RemoteAddr().String())
	}
}
