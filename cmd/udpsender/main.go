package main

import (
	"log"
	"net"
)

// incomplete
func main() {
	conn, err := net.ListenPacket("udp", ":42069")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	dst, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatal(err)
	}

	c, err := net.DialUDP("udp", dst, dst)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

}
