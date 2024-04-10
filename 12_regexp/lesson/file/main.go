package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//1 write
	data := []byte("Hello Bold!")
	text := "Hello Gold!"

	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Unable to write string:", err)
		os.Exit(1)
	}
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Unable to write data:", err)
		os.Exit(1)
	}
	fmt.Println("Done.")

	//2.1 read
	content, err := os.ReadFile("hello.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(string(content))

	//2.2 read
	readFile, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer readFile.Close()

	readData := make([]byte, 64)

	for {
		n, err := readFile.Read(readData)
		if err == io.EOF {
			break
		}
		fmt.Print(string(readData[:n]))
	}

	////3 remove
	//err = os.Remove("hello.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = os.RemoveAll("else")
	//if err != nil {
	//	log.Fatal(err)
	//}

}
