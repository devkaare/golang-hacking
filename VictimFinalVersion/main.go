package main

import (
	"encoding/gob"
	"fmt"
	"golang-hacking/VictimFinalVersion/core/handleConnection"
	"log"
)

type Data struct {
	Name string
	ID   int
	Age  float32
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

	decoder := gob.NewDecoder(connection)

	data := &Data{}

	err = decoder.Decode(data)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("[+] Successfully decoded data")
		fmt.Println(data.Name)
		fmt.Println(data.Age)
		fmt.Println(data.ID)
	}
}
