package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func convertBytesToFile(name string, bytesArr []byte, perm int) {
	err := ioutil.WriteFile(name, bytesArr, os.FileMode(perm))
	if err != nil {
		fmt.Println(err)
		return
	}
}

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

	buff := make([]byte, 1000)
	str := ""
	for /*z := 0; z < 10; z++*/ {
		fmt.Println("-------------------------")
		s := "OK"
		conn.Write([]byte(s))
		//conn.Write([]byte("\r\n\r\n"))
		log.Printf("Send: %s", s)

		n, err := conn.Read(buff)
		if err != nil {
			fmt.Println("END")
			break
		}
		//log.Printf("Receive: %s\n", buff[:n])
		str = str + string(buff[:n])
		//log.Printf("Final packet : %s\n", str)
		//buff = nil

		fmt.Println("-------------------------")
	}
	//convertBytesToFile("packet.jpeg", []byte(str), 0644)
	convertBytesToFile("packet.pdf", []byte(str), 0644)
}
