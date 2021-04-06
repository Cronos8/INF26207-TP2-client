package main

import (
	"fmt"
	"net"
	"os"

	"github.com/Cronos8/INF26207-TP2-client/clientfunc"
	"github.com/Cronos8/INF26207-TP2-client/filebyte"
	"github.com/Cronos8/INF26207-TP2-client/packet"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr file-extension\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example: %s 127.0.0.1:22222 jpeg\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	conn, err := net.Dial("udp4", name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s\n", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	buff := make([]byte, 1024)
	var packetMap map[uint64]int
	packetMap = make(map[uint64]int)
	str := ""

	conn.Write([]byte("Client - CONNEXION OK"))
	fmt.Println("------------------------------------------------")
	fmt.Printf("Connected on : %s\n\n", name)
	fmt.Printf("Local addr : %v\n", conn.LocalAddr().String())
	fmt.Println("------------------------------------------------")

	if clientfunc.NewServerConnexion(conn) == 0 {
		for {
			//fmt.Println("-------------------------")
			n, err := conn.Read(buff)
			if err != nil {
				fmt.Println(err)
				break
			}

			if clientfunc.SendPaquetWithFiability(0.95) == true {
				if n > 0 {

					if string(buff[:n]) == "END" {
						packet.PrintMessage("END FILE TRANSMISSION", packet.PurpleColor, conn.RemoteAddr().String())
						break
					}

					hpacket, buffbody := packet.DecapPacket(buff[:n])

					if packet.IsDuplicatePacket(packetMap, hpacket.HeaderNbPacket) == false {
						packet.PrintMessageWithHeader("Send : PACKAGE RECEIVE", packet.GreenColor, hpacket)
						str = str + string(buffbody)
					} else {
						packet.PrintMessageWithHeader("Send : DUPLICATE PACKAGE", packet.YellowColor, hpacket)
					}
					conn.Write([]byte("PACKAGE RECEIVE"))
					packet.PrintPacket(buff[:n])

				} else {
					continue
				}
			} else {
				packet.PrintMessage("FIABILITY ERROR", packet.RedColor, conn.RemoteAddr().String())
				continue
			}
		}
	}
	filebyte.ConvertBytesToFile("packet."+os.Args[2], []byte(str), 0644)
	filebyte.GetFileByteSignature([]byte(str))
}
