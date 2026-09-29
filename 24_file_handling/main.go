package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		// log the error
		panic(err)
	}

	fileInfo, err := f.Stat()
	if err != nil {
		// log the error
		panic(err)
	}

	fmt.Println("file name:", fileInfo.Name())
	fmt.Println("File or Folder:", fileInfo.IsDir())
	fmt.Println("File size:", fileInfo.Size())
	fmt.Println("File Permission", fileInfo.Mode())
	fmt.Println("File Modified at", fileInfo.ModTime())

	// Read file

	rf, err1 := os.Open("example.txt")
	if err1 != nil {
		panic(err1)
	}

	defer rf.Close()

	buf := make([]byte, 100)

	rf1, err2 := rf.Read(buf)
	if err2 != nil {
		panic(err2)
	}

	for i := range buf {
		fmt.Println("data", rf1, string(buf[i]))
	}

	data, err3 := os.ReadFile("example.txt")
	if err3 != nil {
		panic(err3)
	}

	fmt.Println(string(data))

	// read folder

	dir, err4 := os.Open("../")
	if err4 != nil {
		panic(err4)
	}

	defer dir.Close()

	entries, err5 := dir.ReadDir(-1)

	if err5 != nil {
		panic(err5)
	}

	for _, fi := range entries {
		fmt.Println(fi.Name(), fi.IsDir())
	}

	// create a file

	file, err6 := os.Create("example2.txt")
	if err6 != nil {
		panic(err6)
	}

	defer file.Close()

	file.WriteString("hi golang \n")
	file.WriteString("Nice language \n")

	bytes := []byte("Hello Golang")
	file.Write(bytes)

	// read and write to another file (streaming file)

	sourceFile, err7 := os.Open("example2.txt")

	if err7 != nil {
		panic(err7)
	}

	defer sourceFile.Close()

	destFile, err8 := os.Create("example3.txt")

	if err8 != nil {
		panic(err8)
	}

	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}

		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}
	}

	writer.Flush()

	fmt.Println("Written to new line successfully")

	// Deletion of file

	e := os.Remove("example2.txt")

	if e != nil {
		panic(e)
	}

	fmt.Println("File deleted successfully")
}
