package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	conn, err := net.Dial("udp4", name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("-------------------------")
	fmt.Printf("Connected on : %s\n\n", name)

	s := "Hello Server"
	conn.Write([]byte(s))
	conn.Write([]byte("\r\n\r\n"))
	log.Printf("Send: %s", s)

	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	log.Printf("Receive: %s\n", buff[:n])

	fmt.Println("-------------------------")
}
