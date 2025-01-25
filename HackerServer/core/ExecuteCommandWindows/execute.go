package ExecuteCommandWindows

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

type Command struct {
	CmdOutput string
	CmdErr    string
}

func ExecuteCommandRemotelyWindows(connection net.Conn) (err error) {

	reader := bufio.NewReader(os.Stdin)

	commandloop := true

	for commandloop {

		fmt.Printf(">> ")
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		connection.Write([]byte(command))
		if command == "stop\n" {
			commandloop = false
		}

		cmdStruct := &Command{}

		decoder := gob.NewDecoder(connection)
		err = decoder.Decode(cmdStruct)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(cmdStruct.CmdOutput)
		if cmdStruct.CmdErr != "" {
			fmt.Println(cmdStruct.CmdErr)
		}
	}

	return
}
