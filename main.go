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

	// Récupération des paramètres
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example: %s 127.0.0.1:22222\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	conn, err := net.Dial("udp4", name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s\n", err.Error())
		os.Exit(1)
	}
	// Fermeture de la connexion à la fin
	defer conn.Close()

	buff := make([]byte, 1024)
	var packetMap map[uint64]int
	packetMap = make(map[uint64]int)
	var filename string = ""
	str := ""

	conn.Write([]byte("Client - CONNEXION OK"))
	fmt.Println("------------------------------------------------")
	fmt.Printf("Connected on : %s\n\n", name)
	fmt.Printf("Local addr : %v\n", conn.LocalAddr().String())
	fmt.Println("------------------------------------------------")

	if clientfunc.NewServerConnexion(conn) == 0 {
		// Initialise le nom du fichier
		filename = clientfunc.SetFileName(conn, buff)
		for {
			n, err := conn.Read(buff)
			if err != nil {
				fmt.Println(err)
				break
			}
			// Envoi de la réponse au serveur avec une fiabilité de 95%
			if clientfunc.SendPaquetWithFiability(0.95) == true {
				if n > 0 {

					// Fin de transmission du fichier
					if string(buff[:n]) == "END" {
						packet.PrintMessage("END FILE TRANSMISSION", packet.PurpleColor, conn.RemoteAddr().String())
						break
					}

					// Désencapsulation du paquet reçu
					hpacket, buffbody := packet.DecapPacket(buff[:n])

					// Vérification si le packet n'est pas dupliqué
					if packet.IsDuplicatePacket(packetMap, hpacket.HeaderNbPacket) == false {
						packet.PrintMessageWithHeader("Send : PACKET RECEIVE", packet.GreenColor, hpacket)
						str = str + string(buffbody)
					} else {
						packet.PrintMessageWithHeader("Send : DUPLICATE PACKET", packet.YellowColor, hpacket)
					}
					// Envoi de l'accusé de réception du paquet
					conn.Write([]byte("PACKET RECEIVE"))
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
	// Conversion de la suite d'octets en fichier
	filebyte.ConvertBytesToFile(filename, []byte(str), 0644)
	// Affichage de la signature numérique du fichier
	filebyte.GetFileByteSignature([]byte(str))
}
