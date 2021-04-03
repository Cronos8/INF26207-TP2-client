package main

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
)

func isDuplicatePacket(packetMap map[uint64]int, nbPacket uint64) bool {
	if packetMap[nbPacket] == 1 {
		return true
	}
	packetMap[nbPacket]++
	return false
}

func decapPacket(packet []byte) (uint64, []byte) {
	buffnbpacket := packet[:8]
	buffbody := packet[8:]
	nbPacket := binary.LittleEndian.Uint64(buffnbpacket)

	return nbPacket, buffbody
}

func getFileByteSignature(fileByte []byte) {
	fmt.Printf("File signature : %x\n", sha1.Sum(fileByte))
}

func getByteSignature(packet []byte) {
	log.Printf("Packet signature : %x\n", sha1.Sum(packet))
}

func convertBytesToFile(name string, bytesArr []byte, perm int) {
	err := ioutil.WriteFile(name, bytesArr, os.FileMode(perm))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func sendPaquetWithFiability(fiability float32) bool {
	if rand.Float32() <= fiability {
		return true
	}
	return false
}

//func decapPacket()

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
	fmt.Printf("Local addr : %v || Local network : %v\n", conn.LocalAddr().String(), conn.LocalAddr().Network())
	fmt.Printf("Remote addr : %v || Remote network : %v\n", conn.RemoteAddr().String(), conn.RemoteAddr().Network())
	defer conn.Close()

	buff := make([]byte, 1024)
	buffCo := make([]byte, 1000)
	var packetMap map[uint64]int

	packetMap = make(map[uint64]int)
	str := ""

	conn.Write([]byte("Client - CONNEXION OK"))
	fmt.Println("-------------------------")
	fmt.Printf("Connected on : %s\n\n", name)

	for {
		n, err := conn.Read(buffCo)
		fmt.Println(string(buffCo))
		if err != nil {
			fmt.Println("ERROR")
			continue
		}
		if string(buffCo[:n]) == "Serveur - CONNEXION OK" {
			conn.Write([]byte("READY TO RECEIVE"))
			log.Printf("Send: %s", "READY TO RECEIVE")
			break
		}
	}

	for /*z := 0; z < 10; z++*/ {
		fmt.Println("-------------------------")
		//fmt.Println("dan for")
		n, err := conn.Read(buff)
		//fmt.Println(n)
		//fmt.Println(err)
		//fmt.Println("dan hey")
		if err != nil {
			fmt.Println("ERROR")
			break
		}

		if sendPaquetWithFiability(0.95) == true {
			if n > 0 {
				//fmt.Println("Dans if")
				if string(buff[:n]) == "END" {
					fmt.Println("END")
					break
				}
				nbPacket, buffbody := decapPacket(buff[:n])

				//str = str + string(buff[:n])
				if isDuplicatePacket(packetMap, nbPacket) == false {
					str = str + string(buffbody)
				} else {
					log.Printf("Packet nb : %v DUPLICATE\n", nbPacket)
				}

				conn.Write([]byte("PACKAGE RECEIVE"))
				log.Printf("Send: %s\n", "PACKAGE RECEIVE")
				log.Printf("Package nb : %v\n", nbPacket)
				// log.Println("Packet complet")
				// getByteSignature(buff[:n])
				// log.Println("numero Packet")
				// getByteSignature(buff[:8])
				log.Println("corp du Packet")
				//getByteSignature(buff[8:n])
				getByteSignature(buffbody)

				fmt.Println("-------------------------")

			} else {
				fmt.Println("Dans else")
				continue
			}
		} else {
			log.Println("Fiability Error")
			continue
		}
		//log.Printf("Receive: %s\n", buff[:n])
		//log.Printf("Final packet : %s\n", str)
		//buff = nil

	}
	//convertBytesToFile("packet.jpeg", []byte(str), 0644)
	fmt.Println(packetMap)
	convertBytesToFile("packet.jpg", []byte(str), 0644)
	getFileByteSignature([]byte(str))
}
