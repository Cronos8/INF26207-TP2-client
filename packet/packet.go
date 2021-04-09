package packet

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/Cronos8/INF26207-TP2-client/filebyte"
)

// ColorPrint contient les codes couleurs pour les affichages console
type ColorPrint string

const (
	// BlueColor = bleu
	BlueColor ColorPrint = "\033[34m"
	// RedColor = rouge
	RedColor ColorPrint = "\033[31m"
	// GreenColor = vert
	GreenColor ColorPrint = "\033[32m"
	// ResetColor = couleur d'origine
	ResetColor ColorPrint = "\033[0m"
	// YellowColor = jaune
	YellowColor ColorPrint = "\033[33m"
	// CyanColor = cyan
	CyanColor ColorPrint = "\033[36m"
	// PurpleColor = mauve
	PurpleColor ColorPrint = "\033[35m"
)

// HeaderPacket structure d'une en-tête d'un paquet
type HeaderPacket struct {
	// HeaderIp IPv4 source -> 4 octets
	HeaderIP net.IP
	// HeaderPort Port source -> 4 octets
	HeaderPort int32
	// HeaderNbPacket Numéro du paquet -> 8 octets
	HeaderNbPacket uint64

	// Total : 16 octets
}

// IsDuplicatePacket trouve si un paquet est dupliqué
func IsDuplicatePacket(packetMap map[uint64]int, nbPacket uint64) bool {
	if packetMap[nbPacket] == 1 {
		return true
	}
	packetMap[nbPacket]++
	return false
}

// DecapPacket désencapsule un paquet provenant du serveur
func DecapPacket(packet []byte) (HeaderPacket, []byte) {

	buffnbpacket := packet[:8]
	nbPacket := binary.LittleEndian.Uint64(buffnbpacket)

	buffipacket := packet[8:12]

	buffportpacket := packet[12:16]
	nbPort := binary.LittleEndian.Uint32(buffportpacket)

	hpacket := HeaderPacket{
		net.IP(buffipacket),
		int32(nbPort),
		nbPacket,
	}
	buffbody := packet[16:]

	return hpacket, buffbody
}

// PrintMessage affiche un message client
func PrintMessage(message string, color ColorPrint, ipsource string) {
	fmt.Println(string(color))
	fmt.Println("-----------------------------------------")
	log.Println(message)
	fmt.Println("Serveur addr : " + ipsource)
	fmt.Println("-----------------------------------------")
	fmt.Println(string(ResetColor))
}

// PrintMessageWithHeader affiche un message client en se basant sur l'en-tête du paquet reçu
func PrintMessageWithHeader(message string, color ColorPrint, info HeaderPacket) {
	fmt.Println(string(color))
	fmt.Println("-----------------------------------------")
	log.Println(message)
	fmt.Println("Serveur addr : " + info.HeaderIP.String() + ":" + strconv.Itoa(int(info.HeaderPort)))
	fmt.Println("-----------------------------------------")
	fmt.Println(string(ResetColor))
}

// PrintPacket affiche le numéro du paquet suivi de la signature du paquet entier et de la signature des données
func PrintPacket(p []byte) {

	hpacket, bodyPacket := DecapPacket(p)

	fmt.Printf("\t************************************************\n")
	fmt.Println()

	fmt.Printf("\t[Packet N° : %v]\n", hpacket.HeaderNbPacket)
	fmt.Printf("\tSignature : %x\n", filebyte.GetByteSignature(p))
	fmt.Printf("\tBody Packet - Signature : %x\n", filebyte.GetByteSignature(bodyPacket))

	fmt.Println()
	fmt.Printf("\t************************************************\n")
}
