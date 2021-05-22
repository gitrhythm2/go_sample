package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//basicIo()
	bufferedIo()
}

func basicIo() {
	fmt.Println("basicIo in")
	reader, writer := open()
	copyFile(reader, writer)
}

func bufferedIo() {
	fmt.Println("bufferedIo in")
	r, w := open()
	reader := bufio.NewReader(r)
	writer := bufio.NewWriter(w)
	copyFile(reader, writer)
	writer.Flush() // このFlush()は必要
}

func open() (io.Reader, io.Writer) {
	reader, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	writer, err := os.Create("copy_data.txt")
	if err != nil {
		log.Fatal(err)
	}

	return reader, writer
}

func copyFile(reader io.Reader, writer io.Writer) {
	buf := make([]byte, 32)
	for {
		rlen, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		wlen, err := writer.Write(buf[:rlen])
		if wlen < rlen {
			log.Fatal("Writer error")
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(string(buf[:wlen]))
	}
}
