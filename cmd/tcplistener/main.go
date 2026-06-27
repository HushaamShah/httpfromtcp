package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:42069")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	fmt.Println("Listening on port")

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Connection accepted")
		}

		linech := getLinesChannel(conn)
		for line := range linech {
			fmt.Printf("%s\n", line)
		}

	}

}

func getLinesChannel(f io.ReadCloser) <-chan string {
	buffer := make([]byte, 8)
	ch := make(chan string)

	var currentString string

	go func() {
		defer close(ch)
		defer f.Close()
		for {
			n, err := f.Read(buffer)
			if err != nil {
				if errors.Is(err, io.EOF) {
					ch <- currentString
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				break
			}
			str := string(buffer[:n])
			split := strings.Split(str, "\n")
			for i, v := range split {
				if i == 0 && i == len(split)-1 {
					currentString = currentString + v
				} else if i == len(split)-1 {
					currentString = currentString + v
					// ch <- currentString
				} else {
					currentString = currentString + v
					ch <- strings.TrimSuffix(currentString, "\r")
					currentString = ""
				}
			}
		}
	}()
	return ch
}
