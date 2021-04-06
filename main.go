package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Cronos8/INF26207-TP2-client/clientfunc"
	"github.com/Cronos8/INF26207-TP2-client/filebyte"
	"github.com/Cronos8/INF26207-TP2-client/packet"
)

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

	if clientfunc.NewServerConnexion(conn) == 0 {
		for {
			fmt.Println("-------------------------")
			n, err := conn.Read(buff)
			if err != nil {
				fmt.Println("ERROR")
				break
			}

			if clientfunc.SendPaquetWithFiability(0.95) == true {
				if n > 0 {
					if string(buff[:n]) == "END" {
						fmt.Println("END")
						break
					}
					hpacket, buffbody := packet.DecapPacket(buff[:n])

					if packet.IsDuplicatePacket(packetMap, hpacket.HeaderNbPacket) == false {
						str = str + string(buffbody)
					} else {
						log.Printf("Packet nb : %v DUPLICATE\n", hpacket.HeaderNbPacket)
					}

					conn.Write([]byte("PACKAGE RECEIVE"))
					log.Printf("Send: %s\n", "PACKAGE RECEIVE")
					log.Printf("Package nb : %v\n", hpacket.HeaderNbPacket)
					filebyte.GetByteSignature(buffbody)

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
	filebyte.ConvertBytesToFile("packet.jpeg", []byte(str), 0644)
	filebyte.GetFileByteSignature([]byte(str))
}
