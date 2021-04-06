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

// HeaderPacket header of packet
type HeaderPacket struct {
	HeaderIp       net.IP // 16 byte -> 128 octets
	HeaderPort     int32  // 4 byte -> 32 octets
	HeaderNbPacket uint64 // 8 byte -> 64 octets
	// // 28 bytes au total
}

func isDuplicatePacket(packetMap map[uint64]int, nbPacket uint64) bool {
	if packetMap[nbPacket] == 1 {
		return true
	}
	packetMap[nbPacket]++
	return false
}

func decapPacket2(packet []byte) (uint64, []byte) {
	buffnbpacket := packet[:8]
	buffbody := packet[8:]
	nbPacket := binary.LittleEndian.Uint64(buffnbpacket)

	return nbPacket, buffbody
}

func decapPacket(packet []byte) (HeaderPacket, []byte) {

	buffnbpacket := packet[:8]
	nbPacket := binary.LittleEndian.Uint64(buffnbpacket)

	buffipacket := packet[8:24]

	buffportpacket := packet[24:28]
	nbPort := binary.LittleEndian.Uint32(buffportpacket)

	hpacket := HeaderPacket{
		net.IP(buffipacket),
		int32(nbPort),
		nbPacket,
	}
	buffbody := packet[28:]

	return hpacket, buffbody
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

func newServerConnexion(conn net.Conn) int {

	buffCo := make([]byte, 1000)

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
	return 0
}

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	conn, err := net.Dial("udp4", name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Local addr : %v || Local network : %v\n", conn.LocalAddr().String(), conn.LocalAddr().Network())
	fmt.Printf("Remote addr : %v || Remote network : %v\n", conn.RemoteAddr().String(), conn.RemoteAddr().Network())
	defer conn.Close()

	buff := make([]byte, 1024)
	var packetMap map[uint64]int

	packetMap = make(map[uint64]int)
	str := ""

	conn.Write([]byte("Client - CONNEXION OK"))
	fmt.Println("-------------------------")
	fmt.Printf("Connected on : %s\n\n", name)

	if newServerConnexion(conn) == 0 {
		for {
			fmt.Println("-------------------------")
			n, err := conn.Read(buff)
			if err != nil {
				fmt.Println("ERROR")
				break
			}

			if sendPaquetWithFiability(0.95) == true {
				if n > 0 {
					if string(buff[:n]) == "END" {
						fmt.Println("END")
						break
					}
					hpacket, buffbody := decapPacket(buff[:n])

					if isDuplicatePacket(packetMap, hpacket.HeaderNbPacket) == false {
						str = str + string(buffbody)
					} else {
						log.Printf("Packet nb : %v DUPLICATE\n", hpacket.HeaderNbPacket)
					}

					conn.Write([]byte("PACKAGE RECEIVE"))
					log.Printf("Send: %s\n", "PACKAGE RECEIVE")
					log.Printf("Package nb : %v\n", hpacket.HeaderNbPacket)
					getByteSignature(buffbody)

					fmt.Println("-------------------------")

				} else {
					continue
				}
			} else {
				log.Println("Fiability Error")
				continue
			}
		}
	}
	convertBytesToFile("packet.jpeg", []byte(str), 0644)
	getFileByteSignature([]byte(str))
}
