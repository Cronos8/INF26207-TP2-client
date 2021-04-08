package packet

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/Cronos8/INF26207-TP2-client/filebyte"
)

// ColorPrint colors print
type ColorPrint string

const (
	BlueColor   ColorPrint = "\033[34m"
	RedColor    ColorPrint = "\033[31m"
	GreenColor  ColorPrint = "\033[32m"
	ResetColor  ColorPrint = "\033[0m"
	YellowColor ColorPrint = "\033[33m"
	CyanColor   ColorPrint = "\033[36m"
	PurpleColor ColorPrint = "\033[35m"
)

// HeaderPacket header of packet
type HeaderPacket struct {
	HeaderIp       net.IP // 16 byte -> 128 octets
	HeaderPort     int32  // 4 byte -> 32 octets
	HeaderNbPacket uint64 // 8 byte -> 64 octets
	// // 28 bytes au total
}

func IsDuplicatePacket(packetMap map[uint64]int, nbPacket uint64) bool {
	if packetMap[nbPacket] == 1 {
		return true
	}
	packetMap[nbPacket]++
	return false
}

func DecapPacket2(packet []byte) (uint64, []byte) {
	buffnbpacket := packet[:8]
	buffbody := packet[8:]
	nbPacket := binary.LittleEndian.Uint64(buffnbpacket)

	return nbPacket, buffbody
}

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

func PrintMessage(message string, color ColorPrint, ipsource string) {
	fmt.Println(string(color))
	fmt.Println("-----------------------------------------")
	log.Println(message)
	fmt.Println("Serveur addr : " + ipsource)
	fmt.Println("-----------------------------------------")
	fmt.Println(string(ResetColor))
}

// PrintMessageWithHeader print a message
func PrintMessageWithHeader(message string, color ColorPrint, info HeaderPacket) {
	fmt.Println(string(color))
	fmt.Println("-----------------------------------------")
	log.Println(message)
	fmt.Println("Serveur addr : " + info.HeaderIp.String() + ":" + strconv.Itoa(int(info.HeaderPort)))
	fmt.Println("-----------------------------------------")
	fmt.Println(string(ResetColor))
}

// PrintPacket print a packet
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
